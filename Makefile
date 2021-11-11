run:
	kubectl apply -f ./infra/k8s
stop:
	kubectl delete service auth-srv \
	&& kubectl delete deployment auth-depl \
	&& kubectl delete ingress ingress-service \
	&& kubectl delete service auth-mongo-srv \
	&& kubectl delete deployment auth-mongo-depl \
	&& kubectl delete service client-srv \
	&& kubectl delete deployment client-depl \
	&& kubectl delete service tickets-srv \
	&& kubectl delete deployment tickets-depl \
	&& kubectl delete service tickets-mongo-srv \
	&& kubectl delete deployment tickets-mongo-depl