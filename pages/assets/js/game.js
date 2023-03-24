function uploadGame() {
    var gameName = document.forms["uploadGameTable"]["gameName"].value;
    var gamePrice = document.forms["uploadGameTable"]["gamePrice"].value;
    var gameInfo = document.forms["uploadGameTable"]["gameInfo"].value;
    var gameImg = document.getElementById("gameImg").src;
    var gameGroup = document.forms["uploadGameTable"]["gameGroup"];
    var gameType = ""

    for (var i = 0; i < gameGroup.length; i++) {
        if (gameGroup[i].checked == true) {
            gameType += gameGroup[i].value + ";"
        }
    }

    console.log(gameImg)

    var settings = {
        "url": "http://localhost:8080/user/uploadGame",
        "method": "POST",
        "timeout": 0,
        "headers": {
            "Content-Type": "application/json"
        },
        "data": JSON.stringify({
            "gameName": gameName,
            "gamePrice": gamePrice,
            "gameInfo": gameInfo,
            "gameImg": gameImg,
            "gameType": gameType
        }),
    };

    $.ajax(settings).done(function (response) {
        if (response["status"] == "Accepted") {
            window.location.href = 'result.html';
        } else {
            alert("Fail!");
        }

    });

}
