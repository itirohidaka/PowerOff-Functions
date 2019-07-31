# Power Off a VSI using IBM Cloud Functions
Description: Snippet to Power Off a VSI (Virtual Server) on IBM Cloud using the IBM Cloud Functions.

Requirements:
- Docker CE installed on your Desktop/Laptop.
- Python 2.7 for some local tests.
- I'm using the MacOS X. But you can modify some commands to get it work on Windows.

The following steps will help you to create an Action on the IBM Cloud Functions.
This Action is based on a Python 2 snippet.

### The Easy Way
1. Clone this repo using "git clone" command and access the app folder.
Example:
```
git clone https://github.com/itirohidaka/PowerOff-Functions.git
```
```
cd PowerOff-Functions
```
2. You need to set the Org and Space on ibmcloud cli command. You can use the command bellow to see the orgs and spaces on you account
Example:
```
ibmcloud login
ibmcloud account orgs
ibmcloud account spaces
ibmcloud target -o "WCP Individuals" -s "wcp_uss_itiro"
```
3. Push the zip "package" to IBM Cloud Functions. You can upload my previouly created file (the easy way)
```
ibmcloud fn action create itiroaction01 hello.zip --kind python:2
```

4. Open the IBM Cloud console
```
open https://cloud.ibm.com
```
5. Login to the IBM Cloud Console using your Credentials (username and password)

6. Click on the Three Line Menu (Hamburger Menu) and click on "Functions"

7. Click on "Actions" and check if you Action appear on the list. In my example, I'm using the "itiroaction01" name for the action name.

8. Click on "Trigger" item and click on the "Create" button.

9. Click on "Create Trigger".

10. Click on "Periodic"

11. Type the name on "Trigger Name" field.

12. Select the days and hours that the function will be executed. You can select a pre defined period on "Select pattern" field.

13. In the JSON Payload, type:
```
{
  "username": "<softlayer_username>",
  "key": "<softlayer_api_key>",
  "vsiname":"<name_of_the_vsi>",
  "poweraction":"<power_action>"
}
```
Example:
```
{
  "username": "1234567_itiro@br.ibm.com",
  "key": "ahsdjklfhlajkshdjkfhaljksdhfkahsldhflasjdhfjkashdk",
  "vsiname":"virtualserver01",
  "poweraction":"on"
}
```
OBS: Change <name_of_the_vsi> with the name of the VSI that needs to Powered On/Off.

14. Click on "Create" button.

15. In the next screen, click on "Add" button to associate an action with the Trigger.

16. Click on "Select Existing" Button.

17. Click on "Select an Action" field and Select you action.

18. Click on Add.

19. Your Function is now created with a Trigger and a Action!

You can use the monitor on Function main screen to check if your function is working properly.


### The Hard Way
1. Clone this repo using "git clone" command and access the app folder.
Example:
```
git clone https://github.com/itirohidaka/PowerOff-Functions.git
```
```
cd PowerOff-Functions
```

2. Create the Python virtualenv using a docker command (no need to modify, just copy & paste).
```
docker run --rm -v "$PWD:/tmp" openwhisk/python2action bash -c "cd /tmp && virtualenv virtualenv && source virtualenv/bin/activate && pip install -r requirements.txt"
```

3. Create the zip file with virtualenv folder and \_\_main\_\_.py file.
```
zip -r <zip_file> <virtualenv_folder> <main_file>
```
Example:
```
zip -r hello.zip virtualenv __main__.py
```

4. You need to set the Org and Space on ibmcloud cli command. You can use the command bellow to see the orgs and spaces on you account
Example:
```
ibmcloud login
ibmcloud account orgs
ibmcloud account spaces
ibmcloud target -o "WCP Individuals" -s "wcp_uss_itiro"
```

5. Push the zip "package" to IBM Cloud Functions
```
ibmcloud fn action create <action_name> <zip_file> --kind <runtime>
```
Example:
```
ibmcloud fn action create itiroaction01 hello.zip --kind python:2
```

6. Open the IBM Cloud console
```
open https://cloud.ibm.com
```

7. Login to the IBM Cloud Console using your Credentials (username and password)

8. Click on the Three Line Menu (Hamburger Menu) and click on "Functions"

9. Click on "Actions" and check if you Action appear on the list. In my example, I'm using the "itiroaction01" name for the action name.

10. Click on "Trigger" item and click on the "Create" button.

11. Click on "Create Trigger".

12. Click on "Periodic"

13. Type the name on "Trigger Name" field.

14. Select the days and hours that the function will be executed. You can select a pre defined period on "Select pattern" field.

15. In the JSON Payload, type:
```
{
  "username": "<softlayer_username>",
  "key": "<softlayer_api_key>",
  "vsiname":"<name_of_the_vsi>",
  "poweraction":"<power_action>"
}
```
Example:
```
{
  "username": "1234567_itiro@br.ibm.com",
  "key": "ahsdjklfhlajkshdjkfhaljksdhfkahsldhflasjdhfjkashdk",
  "vsiname":"virtualserver01",
  "poweraction":"on"
}
```
OBS: Change <name_of_the_vsi> with the name of the VSI that needs to Powered On/Off.

16. Click on "Create" button.

17. In the next screen, click on "Add" button to associate an action with the Trigger.

18. Click on "Select Existing" Button.

19. Click on "Select an Action" field and Select you action.

20. Click on Add.

21. Your Function is now created with a Trigger and a Action!

You can use the monitor on Function main screen to check if your function is working properly.

With this snippet you can have an action with several Triggers. Each Trigger with a diffrent VSI name.
