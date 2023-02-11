<template>
    <div class="container">
        <div class="row justify-content-center mt-5">
            <div class="col-lg-4 col-md-4 col-sm-4">
                <div class="card">
                    <div class="card-title text-center">
                        <h2 class="p-3">Login</h2>
                        <div
                            v-if="error"
                            v-on:click="error = false"
                            class="alert alert-danger"
                            role="alert"
                        >
                            Incorrect username or password
                        </div>
                    </div>
                    <div class="card-body">
                        <form v-on:submit.prevent="submit">
                            <div class="mb-3">
                                <label for="email" class="form-label"
                                    >Email:</label
                                >
                                <input
                                    type="email"
                                    name="email"
                                    v-model="email"
                                    class="form-control"
                                />
                            </div>
                            <div class="mb-3">
                                <label for="password" class="form-label"
                                    >Password:</label
                                >
                                <input
                                    type="password"
                                    name="password"
                                    v-model="password"
                                    class="form-control"
                                />
                            </div>
                            <button type="submit" class="btn btn-primary">
                                Submit
                            </button>
                            <p class="float-end">
                                Don't have account
                                <router-link to="/register"
                                    >Sign Up</router-link
                                >
                            </p>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { mapActions } from "vuex"

export default {
    data() {
        return {
            email: "",
            password: "",
            error: false
        }
    },
    methods: {
        ...mapActions(["loginUser"]),
        async submit() {
            const payload = {
                email: this.email,
                password: this.password
            }
            await this.$store
                .dispatch("loginUser", payload)
                .then(() => {
                    this.$router.push("/")
                })
                .catch(() => {
                    this.error = true
                })
        }
    }
}
</script>

<style lang="less" scoped>
.float-end {
    padding-top: 0.5rem;
}
</style>