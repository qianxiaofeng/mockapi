### Usage:
1. go to release page
2. download and unzip source code package
3. download executable into source directory
4. switch between dev and production profile :
    ```bash
    export PROFILE=dev
    ```
5. modify port listen by 
    ```yaml
     port: 8080
    ```
6. run server 
    ```bash
      ./mockapi api run
    ```
   run server daemon
   ```bash
      nohup ./mockapi api run &
   ```
   
