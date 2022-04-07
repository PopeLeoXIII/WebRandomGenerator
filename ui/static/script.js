function myFunction() {
	var nValue = document.getElementById("nInput").value;
	var maxValue = document.getElementById("maxInput").value;
	var genNValue = document.getElementById("genNInput").value;

	alert(nValue + ' ' + maxValue  + ' ' +  genNValue);
}

function myreset() {
	document.getElementById("nInput").setAttribute("value", "") ;
	document.getElementById("maxInput").setAttribute("value", "") ;
	document.getElementById("genNInput").setAttribute("value", "") ;
	//document.getElementById("nInput").value.set("3")

	alert("Данные очищены");
	//document.getElementById('myform').reset()

}