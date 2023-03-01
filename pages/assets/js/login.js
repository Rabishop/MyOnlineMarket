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

  //通过虚拟表单的形式提交post请求，从而实现页面的跳转
  var url = "http://localhost:8080/user/regist";

  $.post(url, {"userAccount":"1235645", "userPassword":"56556"});

}
