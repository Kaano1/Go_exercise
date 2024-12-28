const root = document.getElementById("checkbox-list");

async function getTaskList() {
	var get;

	get = await fetch("http://localhost:9090/showList",{
		method: "GET",
		})
		.then((response) => {
			return response.json();
		})
		.catch((error) => {
			alert(error);
		  });
	console.log(get);
	return get
}

function getShowList() {
	var list;
	
	list = getTaskList();
	console.log(list);
}

function addTaskList() {

}

function removeTaskList() {

}

function completedTaskList() {

}

getShowList()