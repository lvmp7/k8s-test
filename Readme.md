#k8s com minikube

** pre requisitos
 - minukube
 - kubectl

** pre config
 - gerar o secret de password do mysql
   -> kubectl create secret generic mysql-secret --from-literal=password='q1w2e3r4' --from-literal=user='root'

** subindo o ambiente
 - Ngxinx
   -> aplicar configMap.yaml
   -> aplicar deployment.yaml
   -> aplicar service.yaml

 - MySql
   -> aplicar persistent-volume.yaml
   -> aplicar deployment.yaml
   -> aplicar service.yaml


   ---
   Go Web
   - Image
    -> lvmp7/codeeducation-web