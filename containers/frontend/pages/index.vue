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

export default {
    components: {
        Navbar,
        Footer
    },
    middleware: 'user',
}
</script>

<style></style>
