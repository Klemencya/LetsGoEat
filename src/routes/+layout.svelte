<script lang="ts">
    import {base} from "$app/paths";
</script>

<nav>
    <a href="{base}/">Home</a>
</nav>

<slot></slot>

<style>
    nav {
        padding-top: 1%;
        padding-bottom: 1%;
        background-color: black;
        text-decoration-color: white;
        text-decoration-width: auto;
        position:absolute;
        top:0px;
        left:0px;
        width: 100%;
        font-size: 30px;
        text-align: left;
    }

    nav a {
        display: inline;
        padding: 10%
    }

    nav a:visited {
        color: white;
    }
</style>
