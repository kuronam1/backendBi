
/* Обратная связь */

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

/* Студент Homepage */

/*$('#student-brend').on('click', function( event ){
    $.ajax({
        url         : '/studentPanel/menu',
        type        : 'GET',
        dataType    : 'text/html',
        success     : function(data) {
        },
        error: function() {
        }
    });
});
$('#student-journal').on('click', function( event ){
    $.ajax({
        url         : '/studentPanel/journal',
        type        : 'GET',
        dataType    : 'text/html',
        success     : function(data) {
        },
        error: function() {
        }
    });
})
;$('#student-schedule').on('click', function( event ){
    $.ajax({
        url         : '/studentPanel/schedule',
        type        : 'GET',
        dataType    : 'text/html',
        success     : function(data) {
        },
        error: function() {
        }
    });
});*/

/* Parent Homepage */

/*$('#parent-brend').on('click', function( event ){
    $.ajax({
        url         : '/parentPanel/menu',
        type        : 'GET',
        dataType    : 'text/html',
        success     : function(data) {
        },
        error: function() {
        }
    });
});*/
/*$('#parent-journal').on('click', function( event ){
    $.ajax({
        url         : '/parentPanel/journal',
        type        : 'GET',
        dataType    : 'text/html',
        success     : function(data) {
        },
        error: function() {
        }
    });
})*/
/*;$('#parent-schedule').on('click', function( event ){
    $.ajax({
        url         : '/parentPanel/schedule',
        type        : 'GET',
        dataType    : 'text/html',
        success     : function(data) {
        },
        error: function() {
        }
    });
});*/

/* Учительский Homepage */

/*$('#teacher-brend').on('click', function( event ){
    $.ajax({
        url         : '/teacherPanel/menu',
        type        : 'GET',
        dataType    : 'text/html',
        success     : function(data) {
        },
        error: function() {
        }
    });
});
$('#teacher-journal').on('click', function( event ){
    $.ajax({
        url         : '/teacherPanel/journal',
        type        : 'GET',
        dataType    : 'text/html',
        success     : function(data) {
        },
        error: function() {
        }
    });
});
$('#teacher-schedule').on('click', function( event ){
    $.ajax({
        url         : '/teacherPanel/schedule',
        type        : 'GET',
        dataType    : 'text/html',
        success     : function(data) {
        },
        error: function() {
        }
    });
});*/

/* Админ */

/*$('#brend').on('click', function( event ){
    $.ajax({
        url         : '/adminPanel/management',
        type        : 'GET',
        dataType    : 'text/html',
        success     : function(data) {
            window.location.href = '/adminPanel/management';
        },
        error: function() {
        }
    });
});
$('#journal').on('click', function( event ){
    $.ajax({
        url         : '/adminPanel/journal',
        type        : 'GET',
        dataType    : 'text/html',
        success     : function(data) {
        },
        error: function() {
        }
    });
});
$('#schedule').on('click', function( event ){
    $.ajax({
        url         : '/adminPanel/schedule',
        type        : 'GET',
        dataType    : 'text/html',
        success     : function(data) {
        },
        error: function() {
        }
    });
});*/

/* Добавление нового пользователя */

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
            url: '/adminPanel/management/userReg',
            type: 'POST',
            data: JSON.stringify({
                role : roleVale,
                username : nameValue,
                login : loginValue,
                password : passValue
            }),
            dataType: 'json',
            processData: false,
            contentType: 'application/json',
            success: function () {
                alert("Пользователь успешно добавлен")  // Сообщение об успешном добавлении пользователя
            },
            error: function () {
            }
        });
    } else if (document.getElementById("role").value === 'parent') {
        let roleVale = document.getElementById("role").value;
        let nameValue = document.getElementById("name").value;
        let loginValue = document.getElementById("login").value;
        let passValue = document.getElementById("password").value;
        let stndtName = document.getElementById("changeStudent").value;
        $.ajax({
            url: '/adminPanel/management/userReg',
            type: 'POST',
            data: JSON.stringify({
                role : roleVale,
                userName : nameValue,
                login : loginValue,
                password : passValue,
                studentName : stndtName
            }),
            dataType: 'json',
            processData: false,
            contentType: 'application/json',
            success: function () {
                alert("Пользователь успешно добавлен")
            },
            error: function () {
            }
        });
    } else if (document.getElementById("role").value === 'student') {
        let roleVale = document.getElementById("role").value;
        let groupValue = document.getElementById("group").value;
        let nameValue = document.getElementById("name").value;
        let loginValue = document.getElementById("login").value;
        let passValue = document.getElementById("password").value;
        $.ajax({
            url: '/adminPanel/management/userReg',
            type: 'POST',
            data: JSON.stringify({
                role : roleVale,
                groupName : groupValue,
                userName : nameValue,
                login : loginValue,
                password : passValue
            }),
            dataType: 'json',
            processData: false,
            contentType: 'application/json',
            success: function () {
                alert("Пользователь успешно добавлен")
            },
            error: function () {
            }
        });
    }
});

/* Добавление новой группы */

$('#add-group').on( 'click', function( event ){
    let speciality = 'speciality';
    let specialityValue = document.getElementById("groupSpeciality").value;
    let number = 'group';
    let groupValue = document.getElementById("groupNumber").value
    let course = 'course';
    let courseValue = document.getElementById("groupCourse").value
    $.ajax({
        url         : '/adminPanel/management/groupReg', //Проверить на правильность пути к хандлеру
        type        : 'POST',
        data        : JSON.stringify({
                speciality : specialityValue,
                number : groupValue,
                course : courseValue
        }),
        dataType    : 'json',
        processData : false,
        contentType : 'application/json',
        success     : function() {
            alert("Группа успешно добавлена")
        },
        error: function() {
        }
    });
});

/* Админское расписание */

$('#admin-schedule-submit').on( 'click', function( event ){
    if(document.getElementById("group").value === 'false'){
        let teacher = 'teacher';
        let value = document.getElementById("teacher").value;
        $.ajax({
            url         : '/adminPanel/schedule', //Проверить на правильность пути к хандлеру
            type        : 'GET',
            data        : {
                teacher : value
            },
            success     : function(data) {

            },
            error: function() {
            }
        });
    } else if (document.getElementById("teacher").value === 'false'){
        let group = 'group';
        let value = document.getElementById("group").value;
        $.ajax({
            url         : '/adminPanel/schedule', //Проверить на правильность пути к хандлеру
            type        : 'GET',
            data        : {
                group : value
            },
            success     : function(data) {

            },
            error: function() {
            }
        });
    } else {
        $.ajax({
            url         : '/adminPanel/schedule', //Проверить на правильность пути к хандлеру
            type        : 'GET',
            success     : function(data) {

            },
            error: function() {
            }
        });
    }
});

/* Добавление новой оценки */

$('#teacher-new-score').on( 'click', function( event ){
    let disciplineValue = disciplineID;
    let studentName = nameStudent;
    let dateValue = date;
    let gradeValue = document.getElementById("score").value;
    let commentValue = document.getElementById("score-comment").value;
    $.ajax({
        url         : '/teacherPanel/journal',
        type        : 'POST',
        data        : JSON.stringify({
            'studentName' : studentName,
            'disciplineID' : disciplineValue,
            'date' : dateValue,
            'level' : gradeValue,
            'comment' : commentValue
        }),
        dataType    : 'json',
        processData : false,
        contentType : 'application/json',
        success     : function(data) {
            window.location.reload();
        },
        error: function() {
            alert("Ошибка")
        }
    });
});

/* Изменение оценки */

$('#admin-grade-add').on( 'click', function( event ){
    let discipline = 'disciplineID';
    let disciplineValue = disciplineID;
    let name = 'userName';
    let studentName = nameStudent;
    let oldlevel = "oldlevel";
    let oldLevel = scoreLevel;
    let dateName = 'newDate';
    let dateValue = date;
    let gradeID = "gradeID"
    let gradeName = 'newLevel'
    let gradeValue = document.getElementById("admin-grade").value;
    let comment = 'newComment';
    let commentValue = document.getElementById("admin-comment").value;
    $.ajax({
        url         : '/adminPanel/journal/gradesRef',
        type        : 'PATCH',
        data        : JSON.stringify({
            gradeID : levelID,
            name : studentName,
            discipline : disciplineValue,
            oldLevel : oldLevel,
            dateName : dateValue,
            gradeName : gradeValue,
            comment : commentValue
        }),
        dataType    : 'json',
        processData : false,
        contentType : 'application/json',
        success     : function(data) {
            window.location.reload();
        },
        error: function() {
        }
    });
});

/* Добавление новых ДЗ и темы */

$('#sendNewHMandT').on( 'click', function( event ){
    let lessonID = document.getElementById("ThemesHomeWork").attributes[1].value;
    let newHomework = document.getElementById("homeworkNew").value;
    let newThemes = document.getElementById("themesNew").value;
    console.log(lessonID, newHomework, newThemes)
    $.ajax({
        url         : '/teacherPanel/schedule',
        type        : 'POST',
        data        : JSON.stringify({
            'lessonID' : lessonID,
            'subject' : newThemes,
            'homeWork' : newHomework,
        }),
        dataType    : 'json',
        processData : false,
        contentType : 'application/json',
        success     : function(data) {
            window.location.reload();
        },
        error: function() {
        }
    });
});