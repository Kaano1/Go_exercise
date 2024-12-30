var check_list_doc = document.getElementById("checkbox-list");
var fullfill_list_doc = document.getElementById("fullfill-list")
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
		return;
	}

	console.log("get:", get);
	return get
}

async function getShowList() {
	taskMaster = await getTaskList();

	if (!taskMaster || taskMaster[0].completed === undefined) {
		return;
	}

	for (let i = 0; i < taskMaster.length; i++) {
		if (!taskMaster[i].task) {
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

		check_list_doc.appendChild(listItem);
	}
}

function addTask() {
	var listItem = document.createElement("li");
	listItem.id = `check-${data.id}`

	listItem.innerHTML = `
            <input id=${data.id} onclick='completedTaskList(${data.id})' type='checkbox'>
            <strong>${data.task}</strong>
            <button id=remove-${data.id} onclick="removeTaskList(${data.id})" style="color: red; margin-left: 1%;">remove</button>
        `;

	check_list_doc.appendChild(listItem);
}

function addTaskList() {
	data.task = document.getElementById("newTask").value;
	if (data.task.length == 0)
		return;
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
			check_list_doc.innerHTML = ""
			getShowList();
			console.log("Task removed successfully:", data);
		})
		.catch(error => {
			console.error("Error removing task:", error);
		});
}

function completedTaskList(index) {
	console.log("complete index:", index)
	fetch("http://localhost:9090/completedTask",
		{
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify(index),
		}
	)
		.then(response => {
			console.log("Task completed successfully:", response);
		})
}

function observeAllButton() {
	document.getElementById("all-list").hidden = false;
	document.getElementById("completed-list").hidden = true;
	document.getElementById("checkbox-list").innerHTML = "";
	getShowList();
}

async function examine_completed_list() {
	taskMaster = await getTaskList();

	if (!taskMaster || taskMaster[0].completed === undefined) {
		return;
	}

	for (let i = 0; taskMaster[i].task != "" && i < 100; i++) {
		if (taskMaster[i].completed == false)
			continue;
		var listItem = document.createElement("li")
		listItem.id = `${taskMaster[i].id}`

		listItem.innerHTML = `<strong>${taskMaster[i].task}</strong>`;

		fullfill_list_doc.appendChild(listItem);
	}
}

function observCompletedButton() {
	document.getElementById("all-list").hidden = true;
	document.getElementById("completed-list").hidden = false;
	document.getElementById("fullfill-list").innerHTML = "";
	examine_completed_list();
}

getShowList()