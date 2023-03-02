package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/caarlos0/env/v7"
	"github.com/fatih/color"
)

type Config struct {
	BaseUrl         string   `env:"BASE_URL"`
	BearerToken     string   `env:"BEARER_TOKEN"`
	Debug           bool     `env:"DEBUG" envDefault:"false"`
	Deployments     []string `env:"DEPLOYMENTS"`
	DisableOutput   bool     `env:"DISABLE_OUTPUT" envDefault:"false"`
	IgnoreTlsErrors bool     `env:"IGNORE_TLS_ERRORS" envDefault:"false"`
	Namespace       string   `env:"NAMESPACE"`
}

var config = &Config{}

func generateUrl(deployment string) string {
	return fmt.Sprintf("%s/apis/apps/v1/namespaces/%s/deployments/%s", config.BaseUrl, config.Namespace, deployment)
}

func main() {
	color.NoColor = false

	err := env.Parse(config, env.Options{RequiredIfNoDef: true})
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			ResponseHeaderTimeout: 10 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: config.IgnoreTlsErrors,
			},
		},
	}

	// language=json
	reqBody := fmt.Sprintf(`
	{
		"spec": {
			"template": {
				"metadata": {
					"annotations": {
						"kubectl.kubernetes.io/restartedAt": "%s"
					}
				}
			}
		}
	}
	`, time.Now().Format(time.RFC3339))

	ifOutputEnabled(func() {
		fmt.Println("Deployments to redeploy:")

		for _, deployment := range config.Deployments {
			color.Yellow("* %s", deployment)
		}

		fmt.Println("\nStarting to redeploy...")
	})

	errs := Errors{}

	for _, deployment := range config.Deployments {
		req, err2 := http.NewRequest(http.MethodPatch, generateUrl(deployment), strings.NewReader(reqBody))
		if err2 != nil {
			ifOutputEnabled(func() {
				color.Red("⨯ %s", deployment)
			})

			errs = append(errs, &Error{Deployment: deployment, Error: err2})
			continue
		}

		req.Header.Set("authorization", fmt.Sprintf("Bearer %s", config.BearerToken))
		req.Header.Set("content-type", "application/strategic-merge-patch+json")

		res, err2 := httpClient.Do(req)
		if err2 != nil {
			ifOutputEnabled(func() {
				color.Red("⨯ %s", deployment)
			})

			errs = append(errs, &Error{Deployment: deployment, Error: err2})
			continue
		}

		if res.StatusCode == http.StatusOK {
			ifOutputEnabled(func() {
				color.Green("✓ %s", deployment)
			})
		} else {
			ifOutputEnabled(func() {
				color.Red("⨯ %s", deployment)
			})

			resBodyBytes, err3 := io.ReadAll(res.Body)
			if err3 != nil {
				errs = append(errs, &Error{Deployment: deployment, Error: err3})
				continue
			}

			errs = append(
				errs,
				&Error{
					Deployment: deployment,
					Error:      errors.New(fmt.Sprintf("Received a non-OK response:\n%d: %s", res.StatusCode, string(resBodyBytes))),
				},
			)
		}
	}

	if len(errs) > 0 {
		fmt.Printf("\n")

		if config.Debug {
			color.Red("Some errors have occurred while redeploying:")

			for _, err2 := range errs {
				color.Red("\n* %s:\n%s", err2.Deployment, err2.Error)
			}
		} else {
			fmt.Println("There were errors, to see them you can set DEBUG environment variable as 'true'.")
		}

		os.Exit(1)
	}
}

func ifOutputEnabled(f func()) {
	if !config.DisableOutput {
		f()
	}
}
