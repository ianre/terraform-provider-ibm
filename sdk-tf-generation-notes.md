Great, so how do I run the tests?

How the SDK repo got started
https://github.ibm.com/CloudEngineering/go-sdk-template/blob/main/README_FIRST.md


What we did today was figure out the annotations required in the API definition (using test), and you were able to generate the sdk and the terraform that have the new fields. Make sure everything works with the GO SDK using the new version



Command history


 8824  mv openapi-sdkgen-3.102.0.tar.gz openapi-sdkgen
 8825  mv openapi-sdkgen-3.102.0.tar.gz ~/
 8827  mv openapi-sdkgen-3.102.0.tar.gz install-sdk-gen
 8828  tar -xzf openapi-sdkgen-3.102.0.tar.gz
 8829  export PATH=$PATH:/Users/ianre/install-sdk-gen/openapi-sdkgen
 8838  openapi-sdkgen.sh genrate --help
 8839  openapi-sdkgen.sh genrate
 8840  openapi-sdkgen.sh generate
 8841  openapi-sdkgen.sh generate --help
 8842  openapi-sdkgen.sh generate -h
 8843  openapi-sdkgen.sh help
 8845  openapi-sdkgen.sh generate -g ibm-go -i logs-router-service-api.v1.json -o output --genExamples  -v
 8846  openapi-sdkgen.sh generate -g ibm-go -i logs-router-service-api.v1.json -o output --genExamples  -v > build_go.log
 8847  openapi-sdkgen.sh generate -g ibm-terraform -i logs-router-service-api.v1.json -o output-tf
 8849  openapi-sdkgen.sh generate -g ibm-go -i logs-router-service-api.v1.json -o output --genExamples
 8850  openapi-sdkgen.sh generate -g ibm-go -i logs-router-service-api.v1.json -o output --genExamples > examine_with_envvar.log
 8851  openapi-sdkgen.sh generate -g ibm-go -i logs-router-service-api.v1.json -o output --genExamples > ~/install-sdk-gen/openapi-sdkgen/examine_with_envvar.log
 8852  openapi-sdkgen.sh generate -g ibm-go -i logs-router-service-api.v1.json -o output --genExamples
 8853  openapi-sdkgen.sh generate -g ibm-terraform -i logs-router-service-api.v1.json -o ~/install-sdk-gen/openapi-sdkgen
 8855  openapi-sdkgen.sh generate -g ibm-terraform -i logs-router-service-api.v1.json -o ~/install-sdk-gen/openapi-sdkgen/logs
 8859  openapi-sdkgen.sh generate -g ibm-terraform -i logs-router-service-api.v1.json -o ~/install-sdk-gen/openapi-sdkgen/logs > ~/install-sdk-gen/openapi-sdkgen/logs/of-tf-gen.log
 8861  openapi-sdkgen.sh generate -g ibm-terraform -i logs-router-service-api.v1.json -o ~/install-sdk-gen/openapi-sdkgen/logs
 8863  cd openapi-sdkgen
 8899  openapi-sdkgen.sh generate -g ibm-go -t ../../../github.ibm.com/cloud-api-docs/logs-router/logs-router-service-api.v1.json -o test --genExamples
 8904  openapi-sdkgen.sh generate -g ibm-go -i ../../../github.ibm.com/cloud-api-docs/logs-router/logs-router-service-api.v1.json -o test --genExamples
 8905  openapi-sdkgen.sh help generate
 8907  openapi-sdkgen.sh generate -g ibm-terraform  -i ../../../github.ibm.com/cloud-api-docs/logs-router/logs-router-service-api.v1.json -o gentest-tf
 8913  openapi-sdkgen.sh generate -g ibm-go -i ../../../github.ibm.com/cloud-api-docs/logs-router/logs-router-service-api.v1.json -o gentest-go-del --genExamples
 9003  openapi-sdkgen.sh generate -g ibm-terraform  -i ../../../github.ibm.com/cloud-api-docs/logs-router/logs-router-service-api.v1.json -o test-tf



add to ibm/acctest/acctest.go

	IngestionKey = os.Getenv("INGESTION_KEY")
	if IngestionKey == "" {
		IngestionKey = "xxxxxxxxxxxx"
		fmt.Println("[WARN] Set the environment variable INGESTION_KEY for testing Logdna targets, the tests will fail if this is not set")
	}
	return



GNUMakefileTEST?=$$(go list ./... |grep 'routing')
build: fmtcheck [vet]


add the computed fields
add the string to UUID conversion
change the name of the packages
when you do the replace, you need to do go mod tidyreplace github.com/IBM/logs-router-go-sdk => /Users/ianre/go/src/github.com/IBM/logs-router-go-sdk


tf tagsx-ibm-region: https://github.ibm.com/CloudEngineering/openapi-sdkgen/search?q=x-ibm-regionhttps://github.ibm.com/CloudEngineering/openapi-sdkgen/blob/bd651d8274c44013302b5a309ab954bf8af1a4f5/src/test/resources/regional-urls.yamlhttps://github.ibm.com/Observability/pfunk-engineering/issues/2792

Looking for all the other tagshttps://github.ibm.com/search?q=org%3ACloudEngineering+x-resource-name&type=wikis

the pr where we add the annotations in the first place
https://github.ibm.com/cloud-api-docs/logs-router/commit/dc32dfcd26c3abf54b325190320063d6300acbff


UPD8
we are giving up on generating form the api def
we are going to manually fix it again
you pushed the terraform changes to ianre
you also pushed the sdk changes to write-status-1 and 2




make testacc
==> Checking that code complies with gofmt requirements...
TF_ACC=1 go test $(go list ./... |grep 'routing') -v  -timeout 700m
[WARN] Set the environment variable INGESTION_KEY for testing Logdna targets, the tests will fail if this is not set
=== RUN   TestAccIBMLogsRouterTargetsDataSourceBasic
--- PASS: TestAccIBMLogsRouterTargetsDataSourceBasic (73.67s)
=== RUN   TestDataSourceIBMLogsRouterTargetsTargetTypeToMap
--- PASS: TestDataSourceIBMLogsRouterTargetsTargetTypeToMap (0.00s)
=== RUN   TestDataSourceIBMLogsRouterTargetsTargetParametersTypeLogDnaToMap
--- PASS: TestDataSourceIBMLogsRouterTargetsTargetParametersTypeLogDnaToMap (0.00s)
=== RUN   TestDataSourceIBMLogsRouterTargetsTargetTypeLogDnaToMap
--- PASS: TestDataSourceIBMLogsRouterTargetsTargetTypeLogDnaToMap (0.00s)
=== RUN   TestDataSourceIBMLogsRouterTargetsTargetTypeLogsToMap
--- PASS: TestDataSourceIBMLogsRouterTargetsTargetTypeLogsToMap (0.00s)
=== RUN   TestDataSourceIBMLogsRouterTargetsTargetParametersTypeLogsToMap
--- PASS: TestDataSourceIBMLogsRouterTargetsTargetParametersTypeLogsToMap (0.00s)
=== RUN   TestAccIBMLogsRouterTenantsDataSourceBasic
--- PASS: TestAccIBMLogsRouterTenantsDataSourceBasic (74.14s)
=== RUN   TestDataSourceIBMLogsRouterTenantsTenantToMap
--- PASS: TestDataSourceIBMLogsRouterTenantsTenantToMap (0.00s)
=== RUN   TestDataSourceIBMLogsRouterTenantsTargetTypeToMap
--- PASS: TestDataSourceIBMLogsRouterTenantsTargetTypeToMap (0.00s)
=== RUN   TestDataSourceIBMLogsRouterTenantsTargetParametersTypeLogDnaToMap
--- PASS: TestDataSourceIBMLogsRouterTenantsTargetParametersTypeLogDnaToMap (0.00s)
=== RUN   TestDataSourceIBMLogsRouterTenantsTargetTypeLogDnaToMap
--- PASS: TestDataSourceIBMLogsRouterTenantsTargetTypeLogDnaToMap (0.00s)
=== RUN   TestDataSourceIBMLogsRouterTenantsTargetTypeLogsToMap
--- PASS: TestDataSourceIBMLogsRouterTenantsTargetTypeLogsToMap (0.00s)
=== RUN   TestDataSourceIBMLogsRouterTenantsTargetParametersTypeLogsToMap
--- PASS: TestDataSourceIBMLogsRouterTenantsTargetParametersTypeLogsToMap (0.00s)
=== RUN   TestAccIBMLogsRouterTenantBasic
--- PASS: TestAccIBMLogsRouterTenantBasic (75.14s)
=== RUN   TestAccIBMLogsRouterTenantAllArgs
--- PASS: TestAccIBMLogsRouterTenantAllArgs (74.64s)
=== RUN   TestResourceIBMLogsRouterTenantTargetTypeToMap
--- PASS: TestResourceIBMLogsRouterTenantTargetTypeToMap (0.00s)
=== RUN   TestResourceIBMLogsRouterTenantTargetParametersTypeLogDnaToMap
--- PASS: TestResourceIBMLogsRouterTenantTargetParametersTypeLogDnaToMap (0.00s)
=== RUN   TestResourceIBMLogsRouterTenantTargetTypeLogDnaToMap
--- PASS: TestResourceIBMLogsRouterTenantTargetTypeLogDnaToMap (0.00s)
=== RUN   TestResourceIBMLogsRouterTenantTargetTypeLogsToMap
--- PASS: TestResourceIBMLogsRouterTenantTargetTypeLogsToMap (0.00s)
=== RUN   TestResourceIBMLogsRouterTenantTargetParametersTypeLogsToMap
--- PASS: TestResourceIBMLogsRouterTenantTargetParametersTypeLogsToMap (0.00s)
=== RUN   TestResourceIBMLogsRouterTenantMapToTargetTypePrototype
--- PASS: TestResourceIBMLogsRouterTenantMapToTargetTypePrototype (0.00s)
=== RUN   TestResourceIBMLogsRouterTenantMapToTargetParametersTypeLogDnaPrototype
--- PASS: TestResourceIBMLogsRouterTenantMapToTargetParametersTypeLogDnaPrototype (0.00s)
=== RUN   TestResourceIBMLogsRouterTenantMapToTargetTypePrototypeTargetTypeLogDnaPrototype
--- PASS: TestResourceIBMLogsRouterTenantMapToTargetTypePrototypeTargetTypeLogDnaPrototype (0.00s)
=== RUN   TestResourceIBMLogsRouterTenantMapToTargetTypePrototypeTargetTypeLogsPrototype
--- PASS: TestResourceIBMLogsRouterTenantMapToTargetTypePrototypeTargetTypeLogsPrototype (0.00s)
=== RUN   TestResourceIBMLogsRouterTenantMapToTargetParametersTypeLogsPrototype
--- PASS: TestResourceIBMLogsRouterTenantMapToTargetParametersTypeLogsPrototype (0.00s)
PASS
ok  	github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logsrouting	302.113s



This is where we gave up generating the terraform:

Changes to the API DEFhttps://github.ibm.com/cloud-api-docs/logs-router/commit/dc32dfcd26c3abf54b325190320063d6300acbff

CloudEngineering/openapi-sdkgenhttps://github.ibm.com/CloudEngineering/openapi-sdkgen/wiki

wiki searchhttps://github.ibm.com/search?q=org%3ACloudEngineering+x-resource-name&type=wikis

Terraform docs:
https://github.ibm.com/CloudEngineering/openapi-sdkgen/wiki/Design-of-Terraform-generation
https://github.ibm.com/CloudEngineering/openapi-sdkgen/wiki/Terraform-Provider-Generation
https://github.ibm.com/CloudEngineering/openapi-sdkgen/wiki/API-Document-annotations-for-Terraform-&-Ansible-generation#x-terraform-computed-boolean


go-sdk-templatehttps://github.ibm.com/CloudEngineering/go-sdk-template/blob/main/README_FIRST.md#3-modify-selected-files
ibm-cloud-sdk-commonhttps://github.com/IBM/ibm-cloud-sdk-common/blob/main/CONTRIBUTING_go.md#adding-a-new-service



https://github.com/IBM-Cloud/terraform-provider-ibm/issues/5977

They consume us: https://github.com/terraform-ibm-modules/terraform-ibm-cloud-logs/tree/main
BNPP issue: https://github.ibm.com/Observability/pfunk-engineering/issues/2552
Endpoints file: https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/guides/custom-service-endpoints
Docs about our endpoints: https://cloud.ibm.com/docs/logs-router?topic=logs-router-endpoints


The template for where to put the generated files is:https://github.ibm.com/CloudEngineering/terraform-provider-template


