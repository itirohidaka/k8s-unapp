getAll:
	kubectl get nodes
	@echo "\n"
	kubectl get deployments -l app=unapp
	@echo "\n"
	kubectl get pods -l app=unapp
	@echo "\n"
	kubectl get services

getPublicIP:
	ibmcloud cs workers --cluster=freecluster01

getClusterInfo:
	ibmcloud cs cluster-get freecluster01 --showResources

deploy:
	kubectl run unapp --image=registry.ng.bluemix.net/itirohidaka/exemplo
	
deployYaml:
	kubectl apply -f deployment-unapp.yaml

service:
	kubectl expose deployment unapp --port=80 --target-port=80 --type="NodePort"

serviceYaml:
	kubectl apply -f service-unapp.yaml

scale:
	kubectl scale --replicas=5 deployment unapp

rollout:
	kubectl set image deployment/unapp unapp=registry.ng.bluemix.net/itirohidaka/exemplo:v2

rollback:
	kubectl set image deployment/unapp unapp=registry.ng.bluemix.net/itirohidaka/exemplo:latest

rolloutstatus:
	kubectl rollout status deployment/unapp

cleanup:
	kubectl delete -f .
