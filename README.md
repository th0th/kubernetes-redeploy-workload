Deploy kubernetes workloads using kubernetes API.

## Usage

### Running as a GitHub action

```yaml
  - name: Update kubernetes deployments
    uses: th0th/kubernetes-redeploy-workload@v0.1.0
    with:
      base_url: "https://rancher.aperturescience.tld"
      bearer_token: "${{ secrets.BEARER_TOKEN }}"
      debug: "false"
      deployments: "wheatley1,wheatley2"
      disable_output: "false"
      ignore_tls_errors: "false"
      namespace: "namespace"
```

#### Inputs

| Variable          | Required | Default value | Description                                                                                          |
|-------------------|:--------:|---------------|------------------------------------------------------------------------------------------------------|
| base_url          |    ✔     |               | Kubernetes API base url                                                                              |
| bearer_token      |    ✔     |               | Bearer token used for authentication                                                                 |
| debug             |          | 'false'       | Debug flag (useful when something fails)                                                             |
| deployments       |    ✔     |               | Comma separated list of deployment names (e.g. deployment1,deployment2)                              |
| disable_output    |          | 'false'       | Disables outputting to stdout (useful if the logs are public, but you don't want to expose anything) |
| ignore_tls_errors |          | 'false'       | Accept self-signed SSL certificate                                                                   |
| namespace         |    ✔     |               | Kubernetes namespace of the deployment to be updated                                                 |

### Running as a docker container

```shell script
$ docker run --rm -it \
    -e BASE_URL="https://kubernetes.princesscarolyn.com" \
    -e BEARER_TOKEN="bXRwWkNJNklEoyYVdObFlXTmpiM1Z1ZEM5el" \
    -e DEBUG="false" \
    -e DEPLOYMENTS="wheatley1,wheatley2" \
    -e DISABLE_OUTPUT="true" \
    -e IGNORE_TLS_ERRORS="true" \
    -e NAMESPACE="namespace" \
    ghcr.io/th0th/kubernetes-redeploy-workload:0.1.0
```

## Shameless plug

I am an indie hacker, and I am running two services that might be useful for your business. Check them out :)

### WebGazer

[<img alt="WebGazer" src="https://user-images.githubusercontent.com/698079/162474223-f7e819c4-4421-4715-b8a2-819583550036.png" width="256" />](https://www.webgazer.io/?utm_source=github&utm_campaign=rancher-redeploy-workload-readme)

[WebGazer](https://www.webgazer.io/?utm_source=github&utm_campaign=rancher-redeploy-workload-readme) is a monitoring
service that checks your website, cron jobs, or scheduled tasks on a regular basis. It notifies
you with instant alerts in case of a problem. That way, you have peace of mind about the status of your service without
manually checking it.

### PoeticMetric

[<img alt="PoeticMetric" src="https://user-images.githubusercontent.com/698079/162474946-7c4565ba-5097-4a42-8821-d087e6f56a5d.png" width="256" />](https://www.poeticmetric.com/?utm_source=github&utm_campaign=rancher-redeploy-workload-readme)

[PoeticMetric](https://www.poeticmetric.com/?utm_source=github&utm_campaign=rancher-redeploy-workload-readme) is a
privacy-first, regulation-compliant, blazingly fast analytics tool.

No cookies or personal data collection. So you don't have to worry about cookie banners or GDPR, CCPA, and PECR
compliance.

## License

Copyright © 2023, Gökhan Sarı. Released under the [MIT License](LICENSE).
