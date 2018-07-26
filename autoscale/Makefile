CLUSTER?=freecluster01
APP?=autoscale

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

deploy:
	kubectl run php-apache --image=k8s.gcr.io/hpa-example --requests=cpu=200m --expose --port=80

autoscale:
	kubectl autoscale deployment php-apache --cpu-percent=50 --min=1 --max=10
	#while true; do wget -q -O- http://php-apache.default.svc.cluster.local; done
	kubectl run -i --tty load-generator --image=busybox /bin/sh
	
cleanup:
	kubectl delete deployment php-apache
	kubectl delete deployment load-generator
	kubectl delete service php-apache
