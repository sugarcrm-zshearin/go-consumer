# go-kafka

THIS BOTH A BASIC KAFKA PRODUCER AND CONSUMER WRITTEN IN GO.  APPLICATION SETS UP A PRODUCER, PRODUCES A MESSAGE TO THE TOPIC AND THEN SETS UP A CONSUMER AND CONSUMES THAT MESSAGE FROM THE TOPIC

1. Followed instructions [here](https://strimzi.io/quickstarts/) up until right before the send/receive messages section 

2. Modified kafka CRD to add exposed port (so I can communicate with this program as a hardcoded broker) like this:
```
    - name: external
      port: 9094
      tls: false
      type: nodeport
```
3. Get minikube IP
4. Get port by running `kubectl get service my-cluster-kafka-external-bootstrap -n kafka -o=jsonpath='{.spec.ports[0].nodePort}{"\n"}'`
5. Replace broker with `<minikube IP>:<port>` from the last two steps 
6. Run this program (produces and consumes a message)
