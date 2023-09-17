document.addEventListener("DOMContentLoaded", reloadcheckPost); 
function reloadcheckPost(){
    console.log("reload")
    $.ajax({
        type : 'post',
        url : "http://localhost:8080/reloadtask-api",
        contentType: 'application/JSON',
        scriptCharset: 'utf-8'
    })
    .then(
        function(data){
            console.log("status OK");
            console.log(data)
            if (data != null){
                for (let i = 0; i < data.length; i++) {
                    let task = data[i];
                    console.log(task.TaskId + ": " + task.Memo);
                    reloadTasks(task.TaskId, task.Memo)
                }
            }
            
        },
        function(data){
            console.log("status NO");
            window.location.href="http://localhost:8080/login.html"
        }
    );
}

function reloadTasks(taskId, memo){
    document.querySelector('#tasks').innerHTML += `
            <div class="task">
                <span id="taskname">
                    ${memo}
                </span>
                <span id="taskId">${taskId}</span>
                <button class="delete">
                    <i class="far fa-trash-alt"></i>
                </button>
            </div>
        `;

    var current_tasks = document.querySelectorAll(".delete");
    for(var i=0; i<current_tasks.length; i++){
        current_tasks[i].onclick = function(){
            var deleteId = {
                deleteId: this.parentNode.querySelector("#taskId").textContent
            };
            var taskThis = this.parentNode; // thisの値を保存（ajax内だと指す値が変わるため）
            $.ajax({
                type : 'delete',
                url : "http://localhost:8080/deletetask-api",
                data : JSON.stringify(deleteId),
                contentType: 'application/JSON',
                scriptCharset: 'utf-8'
            })
            .then(
                function(data, textStatus, jqXHR){
                    console.log(jqXHR.status)
                    taskThis.remove();
                },
                function(jqXHR, textStatus, errorThrown){
                    console.log(jqXHR.status)
                    if (jqXHR.status >= 500) {
                        alert("server error")
                    }else if (jqXHR.status === 401){
                        alert("Token error, return to login page")
                        window.location.href="http://localhost:8080/login.html"
                    }else if (jqXHR.status >= 400) {
                        alert("request error")
                    }
                }
            );
        }
    }

    var tasks = document.querySelectorAll(".task");
    for(var i=0; i<tasks.length; i++){
        tasks[i].onclick = function(){
            this.classList.toggle('completed');
        }
    }
}