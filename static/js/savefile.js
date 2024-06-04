var files;
$('#admin-save-file').on('change', function(){
    var fileInput= document.getElementById('admin-save-file'); // Начало быдлокода
    var filePath = fileInput.value;
    var allowedExtensions =
        /(.xlsx|.xls)$/i;
    if (!allowedExtensions.exec(filePath)) {
        alert('Неверный формат файла');
        fileInput.files = '';   // Конец быдлокода
    } else files = fileInput.files[0];
});
$('#submit_button').on( 'click', function( event ){
    event.stopPropagation();
    event.preventDefault();
    if( typeof files == 'undefined') return;
    let data = new FormData();
    let groupName = document.getElementById('admin-save-name').value;
    //group
    data.append('group', groupName);
    data.append('file', files);
    $.ajax({
        url         : '/adminPanel/management/scheduleReg',
        type        : 'POST',
        data        : data,
        cache       : false,
        dataType    : 'multipart/form-data',
        processData : false,
        contentType : false,
        success     : function() {
            alert('Файл отправлен')
        }
    });
});

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

/* Перемещение курса на один вперед */

$('#iamsure').on( 'click', function( event ){
    $.ajax({
        url: '/adminPanel/management/updateCourse',
        type: 'POST',
        data: JSON.stringify({
        }),
        dataType: 'json',
        processData: false,
        contentType: 'application/json',
        success: function () {

            alert("Группы перенесены на курс вперед")  // Сообщение об успешном добавлении пользователя
        },
        error: function (data) {
            window.location.reload();
            alert(JSON.parse(data))
        }
    });
});

/* Добавление новой дисциплины */

$('#add-new-discipline').on( 'click', function( event ){
    let teacherValue = document.getElementById("teacherName").value;
    let specialityValue = document.getElementById("groupSpecialityDisc").value;
    let disciplineValue = document.getElementById("disciplineName").value;
    let groupNumberValue = document.getElementById("groupNumberDics").value;
    $.ajax({
        url: '/adminPanel/management/disciplineReg',
        type: 'POST',
        data: JSON.stringify({
            teacherName : teacherValue,
            disciplineName : disciplineValue,
            specialityName : specialityValue,
            course : groupNumberValue
        }),
        dataType: 'json',
        processData: false,
        contentType: 'application/json',
        success: function () {
            window.location.reload();
            alert("Дисциплина успешно добавлена")  // Сообщение об успешном добавлении пользователя
        },
        error: function () {
        }
    });
});

/* Изменение пароля */

$('#change-password').on( 'click', function( event ){
    let loginValue = document.getElementById("login-repass").value;
    let newPassValue = document.getElementById("newPassword").value;
    $.ajax({
        url: '/adminPanel/management/recoverPass',
        type: 'PATCH',
        data: JSON.stringify({
            login : loginValue,
            newPassword : newPassValue
        }),
        dataType: 'json',
        processData: false,
        contentType: 'application/json',
        success: function () {
            alert("Пароль пользователя успешно изменён")  // Сообщение об успешном добавлении пользователя
        },
        error: function () {
        }
    });
});

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
                window.location.reload();
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
                window.location.reload();
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
                window.location.reload();
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
            window.location.reload();
            alert("Группа успешно добавлена")
        },
        error: function() {
        }
    });
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