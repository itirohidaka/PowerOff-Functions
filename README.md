# PowerOff a VSI using IBM Cloud Functions
Description: Snipet to Power Off a VSI (Virtual Server) on IBM Cloud using the IBM Cloud Functions.

Requirements:
- Docker CE installed on your Desktop/Laptop.
- Python 2.7 for some local tests.
- I'm using the MacOS X.

The following steps will help you to create an Action on the IBM Cloud Functions.
This Action is based on a Python 2 snipet.

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
4. Push the zip "package" to IBM Cloud Functions
```
ibmcloud fn action create <action_name> <zip_file> --kind <runtime>
```
Example:
```
ibmcloud fn action create itiroaction01 hello.zip --kind python:2
```
5. Open the IBM Cloud console
```
open https://cloud.ibm.com
```
6. Login to the IBM Cloud Console using your Credentials (username and password)

7. Click on the Three Line Menu (Hamburger Menu) and click on "Functions"

8. Click on "Actions" and check if you Action appear on the list. In my example, I'm using the "itiroaction01" name for the action name.

9. Click on "Trigger" item and click on the "Create" button.

10. Click on "Create Trigger".

11. Click on "Periodic"

12. Type the name on "Trigger Name" field.

13. Select the days and hours that the function will be executed. You can select a pre defined period on "Select pattern" field.

14. In the JSON Payload, type:
```
{
  "vsiname":"<name_of_the_vsi>",
  "poweraction":"<power_action>"
}
```
Example:
```
{
  "vsiname":"virtualserver01",
  "poweraction":"on"
}
```
OBS: Change <name_of_the_vsi> with the name of the VSI that needs to Powered On/Off.

15. Click on "Create" button.

16. In the next screen, click on "Add" button to associate an action with the Trigger.

17. Click on "Select Existing" Button.

18. Click on "Select an Action" field and Select you action.

19. Click on Add.

20. Your Function is now created with a Trigger and a Action!

You can use the monitor on Function main screen to check if your function is working properly.

With this snipet you can have an action with several Triggers. Each Trigger with a diffrent VSI name.
