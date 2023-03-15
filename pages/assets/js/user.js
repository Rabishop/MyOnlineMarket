// set cookie
function setCookie(cname, cvalue, exdays) {
  var d = new Date();
  d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
  var expires = "expires=" + d.toUTCString();
  document.cookie = cname + "=" + cvalue + "; " + expires;
}

function regist() {
  var Account = document.forms["registTable"]["registAccount"].value;
  var Password = document.forms["registTable"]["registPassword"].value;
  var Name = document.forms["registTable"]["registName"].value;

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

  if (Name == null || Name == "") {
    document.getElementById("loginError").innerHTML =
      "Please enter the User Name";
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
      "userName": Name,
    }),
  };

  $.ajax(settings).done(function (response) {
    if (response["status"] == "Account already exists") {
      document.getElementById("loginError").innerHTML =
        "Account already exists";
    } else {
      alert("Sign up Success!");
      window.location.href = 'login.html';
    }

    console.log(response);
  });

}

function login() {
  var Account = document.forms["loginTable"]["loginAccount"].value;
  var Password = document.forms["loginTable"]["loginPassword"].value;

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
    "url": "http://localhost:8080/user/login",
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
    if (response["status"] == "Wrong Username or Password") {
      document.getElementById("loginError").innerHTML =
        "Wrong Username or Password";
    } else {
      alert("Login in Success!");
      window.location.href = 'index.html';
    }

    console.log(response);
  });
}

function logout() {
  var settings = {
    "url": "http://localhost:8080/user/logout",
    "method": "POST",
    "timeout": 0,
    "headers": {
      "Content-Type": "application/json",
    },
  };

  $.ajax(settings).done(function (response) {
    if (response["status"] == "Accepted") {
      alert("Logout success!")
    } else {
      alert("Logout fali!")
    }

    console.log(response);
  });
}

function profileload() {
  var settings = {
    "url": "http://localhost:8080/user/profile",
    "method": "POST",
    "timeout": 0,
    "headers": {
      "Content-Type": "application/json",
    },
  };

  $.ajax(settings).done(function (response) {
    if (response["status"] == "Accepted") {
      document.getElementById("userName").innerHTML =
        response["userName"];
      document.getElementById("userGameCount").innerHTML =
        response["userGameCount"];
      document.getElementById("userPortrait").src = response["userPortrait"];

    } else {
      alert("Please log in first!")
      window.location.href = 'login.html';
    }

    console.log(response);
  });
}

function indexload() {
  var settings = {
    "url": "http://localhost:8080/user/profile",
    "method": "POST",
    "timeout": 0,
    "headers": {
      "Content-Type": "application/json",
    },
  };

  $.ajax(settings).done(function (response) {
    if (response["status"] == "Accepted") {
    } else {
      alert("Please log in first!")
      window.location.href = 'login.html';
    }

    console.log(response);
  });
}

