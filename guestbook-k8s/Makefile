CLUSTER?=freecluster01
APP?=guestbook

getAll:
	kubectl get nodes
	@echo "\n"
	kubectl get deployments
	@echo "\n"
	kubectl get pods -o wide
	@echo "\n"
	kubectl get services -o wide
	@echo "\n"
	kubectl get hpa

getPublicIP:
	ibmcloud cs workers --cluster=${CLUSTER}

getClusterInfo:
	ibmcloud cs cluster-get ${CLUSTER} --showResources

deployYaml:
	kubectl apply -f deployment-redismaster.yaml
	kubectl apply -f deployment-redisslave.yaml
	kubectl apply -f deployment-guestbook.yaml

serviceYaml:
	kubectl apply -f service-redismaster.yaml
	kubectl apply -f service-redisslave.yaml
	kubectl apply -f service-guestbook.yaml

scale:
	kubectl scale --replicas=10 deployment ${APP}

rollout:
	kubectl set image deployment/${APP} ${APP}=ibmcom/guestbook:v2

rollback:
	kubectl set image deployment/${APP} ${APP}=ibmcom/guestbook:v1

status:
	kubectl rollout status deployment/${APP}

cleanup:
	kubectl delete -f .
