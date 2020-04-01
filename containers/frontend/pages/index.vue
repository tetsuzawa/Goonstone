<template>
  <div>
    <header>
      <Navbar/>
    </header>
    <main>
      <div class="container"></div>
    </main>
    <Footer/>
  </div>
</template>

<script>
import Navbar from "../components/Navbar";
import Footer from "../components/Footer";
// import getCookieValue from "../static/js/getCookieValues"

export default {
    components: {
        Navbar,
        Footer
    },
    // methods:{
    //     window:onload = getCookieValue
    // }
}
</script>

<style></style>
