function addPost(){

    var jsonData = {
        memo: document.querySelector('#newtask input').value
    };
    console.log("post json data",jsonData)

    $.ajax({
            type : 'post',
            url : "http://localhost:8080/addtask-api",
            data : JSON.stringify(jsonData),
            contentType: 'application/JSON',
            scriptCharset: 'utf-8'
        })
        .then(
            function(data){
                document.querySelector('#tasks').innerHTML += `
                    <div class="task">
                        <span id="taskname">
                            ${document.querySelector('#newtask input').value}
                        </span>
                        <span id="taskId">${data.taskId}</span>
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
                        $.ajax({
                            type : 'post',
                            url : "http://localhost:8080/deletetask-api",
                            data : JSON.stringify(deleteId),
                            contentType: 'application/JSON',
                            scriptCharset: 'utf-8'
                        })
                        .then(
                            function(data){
                                his.parentNode.remove();
                            },
                            function(data){
                                window.location.href="http://localhost:8080/login.html"
                            }
                        )
                        this.parentNode.remove();
                    }
                }

                var tasks = document.querySelectorAll(".task");
                for(var i=0; i<tasks.length; i++){
                    tasks[i].onclick = function(){
                        this.classList.toggle('completed');
                    }
                }
                
                document.querySelector("#newtask input").value = "";
            },
            function(data){
                console.log("status NO");
                window.location.href="http://localhost:8080/login.html"
            }
        );
}

////////////////////////////////////
document.querySelector('#push').onclick = function(){
    //alert("add Task")
    if(document.querySelector('#newtask input').value.length == 0){
        alert("Please Enter a Task")
    }
    else{
        addPost()
    }
}

function toggleMenu() {
    var menu = document.querySelector(".menu");
    if (menu.style.display === "none") {
      menu.style.display = "block";
    } else {
      menu.style.display = "none";
    }
}

function logout() {
    $.ajax({
        type : 'post',
        url : "http://localhost:8080/logout-api",
        contentType: 'application/JSON',
        scriptCharset: 'utf-8'
    })
    .then(
        function(data){
            console.log("delete status OK");
            window.location.href="http://localhost:8080/login.html"
        },
        function(data){
            console.log("delete status NO");
        }
    );
 }