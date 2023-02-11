<template>
    <div class="container">
        <h2>Edit task</h2>
        <hr />
        <div v-if="task">
            <form @submit.prevent="submit">
                <div class="mb-3">
                    <label for="name" class="form-label">Name</label>
                    <input
                        type="text"
                        name="name"
                        v-model="name"
                        class="form-control"
                    />
                </div>
                <div class="mb-3">
                    <label for="completed" class="form-label">Completed</label>
                    <select class="form-select" v-model="completed">
                        <option value="false">False</option>
                        <option value="true">True</option>
                    </select>
                </div>
                <button type="submit" class="btn btn-primary">Submit</button>
            </form>
        </div>
    </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex"

export default {
    props: ["id"],
    data() {
        return {
            name: "",
            completed: ""
        }
    },
    computed: {
        ...mapGetters({ user: "stateUser", task: "stateTask" })
    },
    methods: {
        ...mapActions(["updateTask", "singleTask"]),
        async submit() {
            let data = {
                id: this.id,
                name: this.name,
                completed: this.completed == "true" ? true : false,
                user_id: this.user.id
            }
            try {
                await this.updateTask(data)
                this.$store.dispatch("userTasks")
                this.$router.push("/dashboard")
            } catch (error) {
                console.error(error)
            }
        },
        async getTask() {
            try {
                await this.singleTask(this.id)
                this.name = this.task.name
                this.completed = this.task.completed
            } catch (error) {
                console.error(error)
            }
        }
    },
    mounted() {
        this.getTask()
    }
}
</script>
