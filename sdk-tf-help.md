The file "/Users/ianre/go/src/github.ibm.com/cloud-api-docs/logs-router/logs-router-service-api.v1.json" contains the api definition for the logs router V1 management api.

We are the Logs router service in IBM Cloud and we want to generate an sdk and terraform using the provided api definition.

In IBM Cloud, we use the repo https://github.ibm.com/CloudEngineering/openapi-sdkgen to generate the sdk and terraform. This repo is available here locally: "/Users/ianre/go/src/github.ibm.com/CloudEngineering/openapi-sdkgen" the wiki is available locally in "/Users/ianre/go/src/github.ibm.com/CloudEngineering/wiki/openapi-sdkgen.wiki"

This respository is owned by a user and will be used to simply PR the changes to the public terraform module repo. 

The repo in "/Users/ianre/go/src/github.ibm.com/Observability/terraform-provider-logs-router" will be the main development repo for new changes to the logs router API. The changes to the actual terraform provider need to be done in this repo but just copied from the repo "/Users/ianre/go/src/github.ibm.com/Observability/terraform-provider-logs-router

<task>
Your task is to read and learn how to use the IBM sdk and terraform generator.
You will then make changes in two repos. The first repo is "/Users/ianre/go/src/github.com/IBM/logs-router-go-sdk" which contains the generated SDK. You can make changes to this repo and I will PR them. The second repo is "/Users/ianre/go/src/github.ibm.com/Observability/terraform-provider-logs-router". You will make changes to this repo so that it can be used as the docs prescribe to test changes to the logs router API definition. Those changes will be tested in that repo and only then will be copied to this repo to raise a PR. You can make any changes to "/Users/ianre/go/src/github.ibm.com/Observability/terraform-provider-logs-router" but please follow the documentation to setup that repo to test new changes to the api definition of service teams. This current repo in "/Users/ianre/go/src/github.com/ianre/terraform-provider-ibm" is only for raining the PRs.
</task>