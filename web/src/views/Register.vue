<template>
    <div class="container">
        <div class="row justify-content-center mt-5">
            <div class="col-lg-4 col-md-4 col-sm-4">
                <div class="card">
                    <div class="card-title text-center">
                        <h3 class="p-3">Register</h3>
                        <div
                            v-if="error"
                            v-on:click="error = false"
                            class="alert alert-danger"
                            role="alert"
                        >
                            User already exists. Please try again.
                        </div>
                    </div>
                    <div class="card-body">
                        <form v-on:submit.prevent="submit">
                            <div class="mb-3">
                                <label for="username" class="form-label"
                                    >Username:</label
                                >
                                <input
                                    type="text"
                                    name="username"
                                    v-model="username"
                                    class="form-control"
                                    :class="{
                                        'is-invalid': $v.username.$error
                                    }"
                                />
                            </div>
                            <div class="mb-3">
                                <label for="email" class="form-label"
                                    >Email:</label
                                >
                                <input
                                    type="text"
                                    name="email"
                                    v-model="email"
                                    class="form-control"
                                    :class="{ 'is-invalid': $v.email.$error }"
                                />
                                <div
                                    v-if="$v.email.$error"
                                    class="invalid-feedback"
                                >
                                    <span v-if="!$v.email.required"
                                        >Email is required</span
                                    >
                                    <span v-if="!$v.email.email"
                                        >Email is invalid</span
                                    >
                                </div>
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
                                    :class="{
                                        'is-invalid': $v.password.$error
                                    }"
                                />
                                <div
                                    v-if="$v.password.$error"
                                    class="invalid-feedback"
                                >
                                    <span v-if="!$v.password.required"
                                        >Password is required</span
                                    >
                                    <span v-if="!$v.password.minLength"
                                        >Password must be at least 6
                                        characters</span
                                    >
                                </div>
                            </div>
                            <div class="mb-3">
                                <label for="confirm_password" class="form-label"
                                    >Confirm password:</label
                                >
                                <input
                                    type="password"
                                    name="password"
                                    v-model="passwordConfirm"
                                    class="form-control"
                                    :class="{
                                        'is-invalid': $v.passwordConfirm.$error
                                    }"
                                />
                                <div
                                    v-if="$v.passwordConfirm.$error"
                                    class="invalid-feedback"
                                >
                                    <span v-if="!$v.passwordConfirm.required"
                                        >Confirm Password is required</span
                                    >
                                    <span
                                        v-else-if="
                                            !$v.passwordConfirm.sameAsPassword
                                        "
                                        >Passwords must match</span
                                    >
                                </div>
                            </div>
                            <button type="submit" class="btn btn-primary">
                                Submit
                            </button>
                            <p class="float-end">
                                Already have account
                                <router-link to="/login">Sign In</router-link>
                            </p>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import axios from "axios"
import { required, email, minLength, sameAs } from "vuelidate/lib/validators"

export default {
    data() {
        return {
            username: "",
            email: "",
            password: "",
            passwordConfirm: "",
            error: false
        }
    },
    validations: {
        username: { required },
        email: { required, email },
        password: { required, minLength: minLength(6) },
        passwordConfirm: { required, sameAsPassword: sameAs("password") }
    },
    methods: {
        async submit() {
            const payload = {
                username: this.username,
                email: this.email,
                password: this.password,
                passwordConfirm: this.passwordConfirm
            }
            this.$v.$touch()
            if (this.$v.$invalid) {
                return
            }
            await axios.post("accounts/register", payload).then((response) => {
                if (!response.data.error) {
                    alert(
                        "You will be redirected to the login page to complete your registration."
                    )
                    this.$router.push("/login")
                    this.error = false
                } else {
                    this.error = true
                }
            })
        }
    }
}
</script>

<style scoped>
.float-end {
    padding-top: 0.5rem;
}
</style>
