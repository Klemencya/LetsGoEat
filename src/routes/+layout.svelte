<script lang="ts">
    import { base } from "$app/paths";
</script>

<nav>
    <a href="{base}/" style="padding-left: 30%">Home</a>
<!--    <a href="{base}/account" style="float:right; padding-right: 30%;"-->
<!--        >My account</a-->
<!--    >-->
</nav>

<slot />

<style>
    @import url("https://fonts.googleapis.com/css2?family=Shrikhand&family=Ubuntu:wght@400;700&display=swap");

    nav {
        padding-top: 1%;
        padding-bottom: 1%;
        background-color: #85d2ac;
        font-size: 20px;
        font-family: "Ubuntu", sans-serif;
        font-weight: 700;
        position: absolute;
        top: 0px;
        left: 0px;
        width: 100%;
        font-size: 24px;
    }

    nav a {
        color: #ffffff;
        text-decoration: none;
        display: inline-block;
        -webkit-transition: all 0.2s ease-in-out;
        transition: all 0.2s ease-in-out;
    }

    nav a:hover {
        color: #9e4eca;
    }
</style>
