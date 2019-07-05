# PowerOff a VSI using IBM Cloud Functions
Snipet to Power Off a VSI (Virtual Server) on IBM Cloud using the Functions.

Requirements:
- Docker CE installed on your Desktop/Laptop.

The following steps will help you to create an Action on the IBM Cloud Functions.
This Action is based on a Python 2 snipet.

1. Clone this repo using "git clone" command and access the app folder.
Example:
git clone https://github.com/itirohidaka/PowerOff-Functions.git
cd PowerOff-Functions

2. Create the Python virtualenv using a docker command (no need to modify, just copy & paste).
docker run --rm -v "$PWD:/tmp" openwhisk/python2action bash -c "cd /tmp && virtualenv virtualenv && source virtualenv/bin/activate && pip install -r requirements.txt"

3. Create the zip file with virtualenv folder and __main__.py file.
Sintaxe:
zip -r <zip_file> <virtualenv_folder> <main_file>

example:
zip -r hello.zip virtualenv __main__.py

4. Push the zip "package" to IBM Cloud Functions
Sintaxe:
ibmcloud fn action create <action_name> <zip_file> --kind <runtime>

Example:
ibmcloud fn action create itiroaction01 hello.zip --kind python:2
