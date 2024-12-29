var root = document.getElementById("checkbox-list");
var taskMaster
var IDs = 0

const data = {
	id: 0,
	task: '',
	completed: '',
};

async function getTaskList() {
	var get;

	try {
		get = await fetch("http://localhost:9090/showList", {
			method: "GET",
		}).then(response => response.json());
	} catch (error) {
		console.log("Error caught:", error);
		return;
	}

	return get
}

async function getShowList() {
	taskMaster = await getTaskList();

	if (!taskMaster || taskMaster[0].completed === undefined) {
		return;
	}

	console.log("hey I'm out");
	for (let i = 0; i < taskMaster.length; i++) {
		if (!taskMaster[i].task) {
			console.log("hey I'm in loop");
			continue;
		}

		var listItem = document.createElement("li");
		listItem.id = `${taskMaster[i].id}`;

		if (taskMaster[i].completed)
			listItem.innerHTML = `            <input id=${taskMaster[i].id} onclick='completedTaskList(${taskMaster[i].id})' type='checkbox' checked>`
		else
			listItem.innerHTML = `            <input id=${taskMaster[i].id} onclick='completedTaskList(${taskMaster[i].id})' type='checkbox'>`
		listItem.innerHTML += `
            <strong>${taskMaster[i].task}</strong>
            <button id=remove-${taskMaster[i].id} onclick="removeTaskList(${taskMaster[i].id})" style="color: red; margin-left: 1%;">remove</button>
        `;

		root.appendChild(listItem);
	}
	console.log(taskMaster);
}

function addTask() {
	var listItem = document.createElement("li");
	listItem.id = `check-${data.id}`

	listItem.innerHTML = `
            <input id=${data.id} onclick='completedTaskList(${data.id})' type='checkbox'>
            <strong>${data.task}</strong>
            <button id=remove-${data.id} onclick="removeTaskList(${data.id})" style="color: red; margin-left: 1%;">remove</button>
        `;

	root.appendChild(listItem);
}

function addTaskList() {
	data.task = document.getElementById("newTask").value;
	data.completed = false;
	data.id = IDs;
	addTask();


	fetch("http://localhost:9090/addTask", {
		method: "POST",
		body: JSON.stringify(data)
	}
	)
		.then((result) => {
			console.log('Success:', result);
		})
		.catch((error) => {
			console.error('Error:', error);
		});
	IDs++;
}

function removeTaskList(id) {
	console.log(id);
	fetch("http://localhost:9090/removeTask", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify(id),
	})
		.then(response => {
			if (!response.ok) {
				throw new Error(`HTTP error! status: ${response.status}`);
			}
			return response.json();
		})
		.then(data => {
			root.innerHTML = ""
			getShowList();
			console.log("Task removed successfully:", data);
		})
		.catch(error => {
			console.error("Error removing task:", error);
		});
}

function completedTaskList(index) {
	console.log(index)
	fetch("http://localhost:9090/completedTask",
		{
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},	
			body: JSON.stringify(index),
		}
	)
}

function observeAllButton() {
	document.getElementById("all-list").hidden = false;
	document.getElementById("completed-list").hidden = true;
	document.getElementById("checkbox-list").innerHTML = "";
	getShowList();
}

function examine_completed_list() {
	for (let i = 0; taskMaster[i].task != "" && i < 100; i++) {
		if (taskMaster[i].completed == false)
			continue;
		var listItem = document.getElementById("fullfill-list").createElement("li")
		listItem.id = `${taskMaster[i].id}`

		listItem.innerHTML = `<strong>${taskMaster[i].task}</strong>`;

		root.appendChild(listItem);
	}
}

function observCompletedButton() {
	document.getElementById("all-list").hidden = true;
	document.getElementById("completed-list").hidden = false;
	document.getElementById("completed-list").innerHTML = "";
	examine_completed_list();
}

getShowList()