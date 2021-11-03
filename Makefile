run:
	kubectl apply -f ./infra/k8s
stop:
	kubectl delete service auth-srv \
	&& kubectl delete deployment auth-depl \
	&& kubectl delete ingress ingress-service