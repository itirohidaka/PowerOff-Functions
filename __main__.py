import SoftLayer
import json

def main(args):
    # Extract the id (VSI ID) from JSON payload. JSON payload received by the Trigger
    name = args.get("payload")
    namejson = json.loads(name)

    """
    # retrive the variables from JSON Payload
    """
    virtualGuestName = namejson["vsiname"]
    print("VSI Name: " + virtualGuestName)
    power = namejson["poweraction"]
    print("Power Action: " + power)
    ibmcloud_iaas_user = namejson["username"]
    print("Username: " + ibmcloud_iaas_user)
    ibmcloud_iaas_key = namejson["key"]
    print("API Key: " + ibmcloud_iaas_key)

    """
    # Your SoftLayer API username and key.
    # Generate an API key at the SoftLayer Customer Portal:
    # control.softlayer.com
    """
    username = 'ibmcloud_iaas_user'
    key = 'ibmcloud_iaas_key'

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

    if power == "off":
        try:
            # Power off the virtual guest
            virtualMachines = client['SoftLayer_Virtual_Guest'].powerOff(id=virtualGuestId)
            print (virtualGuestName + " powered off")
            return { "Status": "OK" }
        except SoftLayer.SoftLayerAPIError as e:
            print("Unable to power on/off the virtual guest"
                  % (e.faultCode, e.faultString))
    elif power == "on":
        try:
            # Power off the virtual guest
            virtualMachines = client['SoftLayer_Virtual_Guest'].powerOn(id=virtualGuestId)
            print (virtualGuestName + " powered off")
            return { "Status": "OK" }
        except SoftLayer.SoftLayerAPIError as e:
            print("Unable to power on/off the virtual guest"
                  % (e.faultCode, e.faultString))
