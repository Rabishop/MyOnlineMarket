function uploadGame() {
    var gameName = document.forms["uploadGameTable"]["gameName"].value;
    var gamePrice = document.forms["uploadGameTable"]["gamePrice"].value;
    var gameGroup = document.forms["uploadGameTable"]["gameGroup"];
    var gameType = ""

    for (var i = 0; i < gameGroup.length; i++) {
        if (gameGroup[i].checked == true) {
            gameType += gameGroup[i].value + ";"
        }
    }

    console.log(gameType)

    // var settings = {
    //     "url": "http://localhost:8080/user/regist",
    //     "method": "POST",
    //     "timeout": 0,
    //     "headers": {
    //         "Content-Type": "application/json"
    //     },
    //     "data": JSON.stringify({
    //         "userAccount": Account,
    //         "userPassword": Password,
    //         "userName": Name,
    //     }),
    // };

    // $.ajax(settings).done(function (response) {
    //     if (response["status"] == "Account already exists") {
    //         document.getElementById("loginError").innerHTML =
    //             "Account already exists";
    //     } else {
    //         alert("Sign up Success!");
    //         window.location.href = 'login.html';
    //     }

    //     console.log(response);
    // });

}
