new Vue({
    el: 'body',
    data: {
        tasks: [],
        newTask: {}
    },
    created: function () {
        this.$http.get('/tasks').then(function (response) {
            this.tasks = response.data ? response.data : []
        })
    },
    methods: {
        createTask: function () {
            if (!$.trim(this.newTask.name)) {
                this.newTask = {}
                return
            }
            this.$http.put('/tasks', this.newTask).success(function (response) {
                this.newTask.id = response.created
                this.tasks.push(this.newTask)
                console.log("Task created!")
                console.log(this.newTask)
                this.newTask = {}
            }).error(function(error) {
                console.log(error)
            })
        },
        deleteTask: function (index) {
            this.$http.delete('/tasks/'+this.tasks[index].id).success(function (response) {
                this.tasks.splice(index, 1)
                console.log("Task deleted!")
            }).error(function (error) {
                console.log(error)
            })
        }
    }
})