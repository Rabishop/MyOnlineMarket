function login() {
  var Account = document.forms["login"]["loginAccount"].value;
  var Password = document.forms["login"]["loginPassword"].value;

  if (Account == null || Account == "") {
    document.getElementById("loginError").innerHTML =
      "Please enter the Account";
    return false;
  }

  if (Password == null || Password == "") {
    document.getElementById("loginError").innerHTML =
      "Please enter the Password";
    return false;
  }


  var settings = {
    "url": "http://localhost:8080/user/regist",
    "method": "POST",
    "timeout": 0,
    "headers": {
      "Content-Type": "application/json"
    },
    "data": JSON.stringify({
      "userAccount": Account,
      "userPassword": Password,
    }),
  };

  $.ajax(settings).done(function (response) {
    console.log(response);
  });

}