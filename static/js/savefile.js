var files;
$('#files').on('change', function(){
    var fileInput = document.getElementById('files'); // Начало быдлокода
    var filePath = fileInput.value;
    var allowedExtensions =
        /(\.xlsx|\.xls)$/i;
    if (!allowedExtensions.exec(filePath)) {
        alert('Неверный формат файла');
        fileInput.value = '';   // Конец быдлокода
    } else files = this.files;
});
$('#submit_button').on( 'click', function( event ){

    event.stopPropagation();
    event.preventDefault();
    if( typeof files == 'undefined' ) return;
    var data = new FormData();
    $.each( files, function( key, value ){
        console.log(key);
        console.log(value);
        data.append( key, value );
    });
    data.append( 'my_file_upload', 1 );
    $.ajax({
        url         : '',
        type        : 'POST',
        data        : data,
        cache       : false,
        dataType    : 'json',
        processData : false,
        contentType : false,
        success     : function() {
            alert('Файл отправлен')
        },
        error: function() {
            alert('Ошибка передачи файла');
        }
    });

});
$('#enter').on( 'click', function( event ){
        let login = 'login';
        let loginValue = document.getElementById("index-login").value;
        let password = 'password';
        let passValue = document.getElementById("index-pass").value;
    $.ajax({
        url         : '/auth',
        type        : 'POST',
        data        : JSON.stringify({
            login : loginValue,
            password : passValue
        }),
        dataType    : 'json',
        processData : false,
        contentType : 'application/json',
        success     : function(data) {
            if(data.redirectUrl) {
                window.location.href = data.redirectUrl;
            }
            },
        error: function() {

        }
    });
});
$('#newsletter-submit').on( 'click', function( event ){
    let fio = 'fio';
    let fioVale = document.getElementById("fio").value;
    let email = 'email';
    let emailVale = document.getElementById("email").value;
    let message = 'message';
    let messageValue = document.getElementById("message-news").value;
    $.ajax({
        url         : '/feedback',
        type        : 'POST',
        data        : JSON.stringify({
            fio : fioVale,
            email : emailVale,
            message : messageValue
        }),
        dataType    : 'json',
        processData : false,
        contentType : 'application/json',
        success     : function() {
        },
        error: function() {
        }
    });
});
$('#student-brend').on('click', function( event ){
    $.ajax({
        url         : '/studentPanel/menu',
        type        : 'GET',
        dataType    : 'text/html',
        success     : function(data) {
            alert($(data).html());
        },
        error: function() {
            alert("Проблемы");
        }
    });
});
$('#student-journal').on('click', function( event ){
    $.ajax({
        url         : '/studentPanel/journal',
        type        : 'GET',
        dataType    : 'text/html',
        success     : function(data) {
            alert($(data).html());
        },
        error: function() {
            alert("Проблемы");
        }
    });
})
;$('#student-schedule').on('click', function( event ){
    $.ajax({
        url         : '/studentPanel/schedule',
        type        : 'GET',
        dataType    : 'text/html',
        success     : function(data) {
            alert($(data).html());
        },
        error: function() {
            alert("Проблемы");
        }
    });
});
$('#teacher-brend').on('click', function( event ){
    $.ajax({
        url         : '/teacherPanel/menu',
        type        : 'GET',
        dataType    : 'text/html',
        success     : function(data) {
            alert($(data).html());
        },
        error: function() {
            alert("Проблемы");
        }
    });
});
$('#teacher-journal').on('click', function( event ){
    $.ajax({
        url         : '/teacherPanel/journal',
        type        : 'GET',
        dataType    : 'text/html',
        success     : function(data) {
            alert($(data).html());
        },
        error: function() {
            alert("Проблемы");
        }
    });
});
$('#teacher-schedule').on('click', function( event ){
    $.ajax({
        url         : '/teacherPanel/schedule',
        type        : 'GET',
        dataType    : 'text/html',
        success     : function(data) {
            alert($(data).html());
        },
        error: function() {
            alert("Проблемы");
        }
    });
});
$('#brend').on('click', function( event ){
    $.ajax({
        url         : '/adminPanel/management',
        type        : 'GET',
        dataType    : 'text/html',
        success     : function(data) {
            alert($(data).html());
        },
        error: function() {
            alert("Проблемы");
        }
    });
});
$('#journal').on('click', function( event ){
    $.ajax({
        url         : '/adminPanel/journal',
        type        : 'GET',
        dataType    : 'text/html',
        success     : function(data) {
            alert($(data).html());
        },
        error: function() {
            alert("Проблемы");
        }
    });
});
$('#schedule').on('click', function( event ){
    $.ajax({
        url         : '/adminPanel/schedule',
        type        : 'GET',
        dataType    : 'text/html',
        success     : function(data) {
            alert($(data).html());
        },
        error: function() {
            alert("Проблемы");
        }
    });
});
$('#new-user-reg').on( 'click', function( event ){
    if(document.getElementById("role").value === 'teacher'){
        let role = 'role';
        let roleVale = document.getElementById("role").value;
        let name = 'name';
        let nameValue = document.getElementById("name").value;
        let login = 'login';
        let loginValue = document.getElementById("login").value;
        let password = 'password';
        let passValue = document.getElementById("password").value;
    $.ajax({
        url: '/adminPanel/management/userRegistration',
        type: 'POST',
        data: JSON.stringify({
            role : roleVale,
            name : nameValue,
            login : loginValue,
            password : passValue
        }),
        dataType: 'json',
        processData: false,
        contentType: 'application/json',
        success: function () {
            //Пока это оставим пустым, потом если надо, что-то придумаем
        },
        error: function () {
        }
    });
    } else {
        let role = 'role';
        let roleVale = document.getElementById("role").value;
        let group = 'group';
        let groupValue = document.getElementById("group").value;
        let name = 'name';
        let nameValue = document.getElementById("name").value;
        let login = 'login';
        let loginValue = document.getElementById("login").value;
        let password = 'password';
        let passValue = document.getElementById("password").value;
        $.ajax({
            url: '/adminPanel/management/userRegistration',
            type: 'POST',
            data: JSON.stringify({
                role : roleVale,
                group : groupValue,
                name : nameValue,
                login : loginValue,
                password : passValue
            }),
            dataType: 'json',
            processData: false,
            contentType: 'application/json',
            success: function () {
                //Пока это оставим пустым, потом если надо, что-то придумаем
            },
            error: function () {
            }
        });
    }
});
$('#add-group').on( 'click', function( event ){
    let speciality = 'speciality';
    let specialityValue = document.getElementById("role").value;
    let group = 'group';
    let groupValue = document.getElementById("groupNumber").value
    let course = 'course';
    let courseValue = document.getElementById("groupCourse").value
    $.ajax({
        url         : '/adminPanel/managment/groupReg', //Проверить на правильность пути к хандлеру
        type        : 'POST',
        data        : JSON.stringify({
                speciality : specialityValue,
                group : groupValue,
                course : courseValue
        }),
        dataType    : 'json',
        processData : false,
        contentType : 'application/json',
        success     : function() {
            //Пока это оставим пустым, потом если надо, что-то придумаем
        },
        error: function() {
        }
    });
});
$('#admin-schedule-submit').on( 'click', function( event ){
    if(document.getElementById("group").value === 'false'){
        let group = 'teacher';
        let value = document.getElementById("teacher").value;
        $.ajax({
            url         : '/adminPanel/schedule', //Проверить на правильность пути к хандлеру
            type        : 'GET',
            data        :JSON.stringify( {
                group : value
            }),
            success     : function(data) {
                alert(data);
            },
            error: function() {
            }
        });
    } else if (document.getElementById("teacher").value === 'false'){
        let teacher = 'group';
        let value = document.getElementById("group").value;
        $.ajax({
            url         : '/adminPanel/schedule', //Проверить на правильность пути к хандлеру
            type        : 'GET',
            data        : JSON.stringify({
                teacher : value
            }),
            success     : function(data) {
                alert(data);
            },
            error: function() {
            }
        });
    } else {
        $.ajax({
            url         : '/adminPanel/schedule', //Проверить на правильность пути к хандлеру
            type        : 'GET',
            success     : function(data) {
                alert(data);
            },
            error: function() {
            }
        });
    }
});
$('#cancel-button').on( 'click', function( event ){
    $.ajax({
        url         : '/logout',
        type        : 'POST',
        date : JSON.stringify({}),
        dataType    : 'json',
        processData : false,
        contentType : 'application/json',
        success     : function() {
            //Пока это оставим пустым, потом если надо, что-то придумаем
        },
        error: function() {
            alert('logout');
        }
    });
});
$('#admin-get-journale-group').on( 'click', function( event ){
    let group = 'group';
    let groupValue = document.getElementById("admin-group").value;
    let discipline = 'discipline';
    let disciplineValue = document.getElementById("discipline").value;
    $.ajax({
        url         : '/adminPanel/schedule',
        type        : 'GET',
        data        : JSON.stringify({
                group : groupValue,
                discipline : disciplineValue
        }),
        success     : function(data) {
            alert(data);
        },
        error: function() {
        }
    });
});
$('#teacher-get-journal').on( 'click', function( event ){
    let group = 'group';
    let groupValue = document.getElementById("teacher-group").value;
    let discipline = 'discipline';
    let disciplineValue = document.getElementById("teacher-discipline").value;
    $.ajax({
        url         : '/teacherPanel/schedule',
        type        : 'GET',
        data        : JSON.stringify({
            discipline : disciplineValue,
            group : groupValue
        }),
        success     : function(data) {
            alert(data);
        },
        error: function() {
        }
    });
});
$('#teacher-new-score').on( 'click', function( event ){
    let discipline = 'discipline';
    let disciplineValue = document.getElementById("teacher-discipline").value;
    let name = 'name';
    let studentName = nameStudent;
    let dateName = 'date';
    let dateValue = date;
    let gradeName = 'grade'
    let gradeValue = document.getElementById("score").value;
    let comment = 'comment';
    let commentValue = document.getElementById("score-comment").value;
    $.ajax({
        url         : 'teacherPanel/journal',
        type        : 'POST',
        data        : JSON.stringify({
            name : studentName,
            discipline : disciplineValue,
            dateName : dateValue,
            gradeName : gradeValue,
            comment : commentValue
        }),
        dataType    : 'json',
        processData : false,
        contentType : 'application/json',
        success     : function(data) {

        },
        error: function() {
            alert('Ошибка');
        }
    });
});
$('#admin-grade-add').on( 'click', function( event ){
    let discipline = 'discipline';
    let disciplineValue = document.getElementById("discipline").value;
    let name = 'name';
    let studentName = nameStudent;
    let dateName = 'date';
    let dateValue = date;
    let gradeName = 'grade'
    let gradeValue = document.getElementById("admin-grade").value;
    alert(gradeValue);
    let comment = 'comment';
    let commentValue = document.getElementById("admin-comment").value;
    $.ajax({
        url         : 'teacherPanel/journal',
        type        : 'POST',
        data        : JSON.stringify({
            name : studentName,
            discipline : disciplineValue,
            dateName : dateValue,
            gradeName : gradeValue,
            comment : commentValue
        }),
        dataType    : 'json',
        processData : false,
        contentType : 'application/json',
        success     : function(data) {

        },
        error: function() {
            alert('Ошибка');
        }
    });
});
//     encodeURIComponent( {
//             discipline : disciplineValue,
//             group : groupValue
//         }),