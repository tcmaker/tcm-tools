import json, requests, sys

authToken = sys.argv[1]
personsURL = "https://members.tcmaker.org/api/v1/persons/?format=json"
headers = {"Authorization": "Bearer " + authToken}

f = open("members_export.csv", "a")
f.write("Email Address,First Name,Last Name,Address,Phone\n")

while True: 
    print(personsURL)
    pageData = requests.get(personsURL, headers=headers).json()

    for person in pageData["results"]:
        if len(person["administered_households"]) > 0 and person["administered_households"][0]["status"] == "active":
            street2 = " " + person["address_street2"] if person["address_street2"] else ""
            f.write(person["email"] + "," + 
                person["given_name"] + "," + 
                person["family_name"] + "," +
                "\"" + person["address_street1"] + street2 + " " + person["address_city"] + "," + person["address_state"] + " " + person["address_zip"] + "\"," +
                person["phone_number"] + "\n")

        personsURL = pageData["next"]

        if personsURL == None:
            break

f.close()
print("Complete")