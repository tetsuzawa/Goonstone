<template>
  <footer class="footer">
    <button v-if="hasLoggedIn" @click="logout" class="button button--link">Logout</button>
    <nuxt-link to="/login" v-else class="button button--link">Login / Register</nuxt-link>
  </footer>
</template>

<script>
    export default {
        name: "Footer",
        methods: {
            async logout() {
                await this.$store.dispatch('auth/logout')
                this.$router.push('/login')
            },
        },
        computed: {
            hasLoggedIn() {
                return this.$store.getters['auth/hasLoggedIn']
            },
        }
    }
</script>

<style scoped>

</style>
