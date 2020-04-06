<template>
  <div class="container-small">
    <ul class="tab">
      <li class="tab-item" :class="{'tab-item-active': tab === 1}" @click="tab = 1">Login</li>
      <li class="tab-item" :class="{'tab-item-active': tab === 2}" @click="tab = 2">Register</li>
    </ul>
    <div class="panel" v-show="tab === 1">
      <form class="form" @submit.prevent="login">
        <label for="login-email">Email</label>
        <input type="text" id="login-email" v-model="loginForm.email" class="form-item">
        <label for="login-password">Password</label>
        <input type="password" id="login-password" v-model="loginForm.password" class="form-item">
        <div class="form-button">
          <button type="submit" class="button button-inverse">Login</button>
        </div>
      </form>
    </div>
    <div class="panel" v-show="tab === 2">
      <form class="form" @submit.prevent="register">
        <label for="username">Name</label>
        <input type="text" id="username" v-model="registerForm.name" class="form-item">
        <label for="email">Email</label>
        <input type="text" id="email" v-model="registerForm.email" class="form-item">
        <label for="password">Password</label>
        <input type="password" id="password" v-model="registerForm.password" class="form-item">
        <label for="password-confirmation">Password (confirm)</label>
        <input type="password" id="password-confirmation" v-model="registerForm.password_confirmation"
               class="form-item">
        <div class="form-button">
          <button type="submit" class="button button-inverse">register</button>
        </div>
      </form>
    </div>
    <div class="panel" v-show="tab === 2">Register Form</div>
  </div>
</template>

<script>
    import Footer from "../components/Footer";
    import Index from "./index";
    import authenticated from "../middleware/authenticated";

    export default {
        name: 'Login',
        components: {Index, Footer},
        middleware: authenticated,
        data() {
            return {
                tab: 1,
                loginForm: {
                    email: '',
                    password: ''
                },
                registerForm: {
                    name: '',
                    email: '',
                    password: '',
                    password_confirmation: ''
                }
            }
        },
        methods: {
            async register() {
                await this.$store.dispatch('auth/register', this.registerForm);
                if (this.apiStatusHasError) {
                    this.$router.push('/')
                }
            },
            async login() {
                await this.$store.dispatch('auth/login', this.loginForm);
                if (this.apiStatusHasError) {
                    this.$router.push('/')
                }
            },
        },
        computed: {
            apiStatusHasError() {
                return this.$store.state.auth.apiStatusHasError
            }
        }
    }
</script>

<style scoped></style>
