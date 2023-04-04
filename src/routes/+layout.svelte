<script lang="ts">
    import {base} from "$app/paths";
</script>

<nav>
    <a href="{base}/" style="padding-left: 10%">Home</a>
    <a href="{base}/account" style="float:right; padding-right: 10%; text-decoration-color: white; color: white">My account</a>
</nav>

<slot></slot>

<style>
    nav {
        padding-top: 1%;
        padding-bottom: 1%;
        background-color: #85d2ac;
        text-decoration-color: white;
        text-decoration-width: none;
        position:absolute;
        top:0px;
        left:0px;
        width: 100%;
        font-size: 20px;
        text-align: left;
    }

    nav a {
        display: inline;
    }

    nav a:visited {
        color: white;
    }
</style>
