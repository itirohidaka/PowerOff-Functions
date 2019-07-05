import SoftLayer
import json

def main(args):
    greeting = "Hello " + "!"

    # Extract the id (VSI ID) from JSON payload. JSON payload received by the Trigger
    name = args.get("payload")
    namejson = json.loads(name)

    """
    # The name of the machine you wish to power off.
    # JSON Template to be included in the Trigger: { "vsiname":"<VSIname>"}.
    # Change the <VSIname> with the VSI hostname
    """
    virtualGuestName = namejson["vsiname"]
    print("VSI Name: " + virtualGuestName)

    """
    # Your SoftLayer API username and key.
    # Generate an API key at the SoftLayer Customer Portal:
    # control.softlayer.com
    """
    username = '1713407_itiro@br.ibm.com'
    key = '48eccbe805a03d71558cf8fb37b433172acd48bcc194cc7c14965d3b5484853a'

    # Declare a new API service object
    client = SoftLayer.Client(username=username, api_key=key)

    try:
        # Getting all virtual guest that the account has:
        virtualGuests = client['SoftLayer_Account'].getVirtualGuests()
    except SoftLayer.SoftLayerAPIError as e:
        print("Unable to retrieve hardware. "
              % (e.faultCode, e.faultString))

    # Looking for the virtual guest
    virtualGuestId = ''
    for virtualGuest in virtualGuests:
        if virtualGuest['hostname'] == virtualGuestName:
            virtualGuestId = virtualGuest['id']

    print("VSI ID:" + str(virtualGuestId))

    try:
        # Power off the virtual guest
        virtualMachines = client['SoftLayer_Virtual_Guest'].powerOff(id=virtualGuestId)
        print (virtualGuestName + " powered off")
        return { "Status": "powered off" }
    except SoftLayer.SoftLayerAPIError as e:
        print("Unable to power off the virtual guest"
              % (e.faultCode, e.faultString))