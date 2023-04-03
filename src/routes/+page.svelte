<script lang="ts">
    import {base} from "$app/paths";
</script>

<div class="form">
    <h1>Let's Go Eat</h1>
    <nav>
        <a href="{base}/login"><button>Log in</button></a>
        <a href="{base}/registration"><button>Registration</button></a>
    </nav>
</div>

<style>
    .form {
        text-align: center;
        margin-right: 30%;
        margin-left: 30%;
        margin-top: 10%;
        border: #9e4eca;
    }

    div h1 {
        font-size: 40px;
    }

    div nav {
        display: flex;
        flex-direction: column;
    }

    nav button {
        margin-top: 20px;
        padding: 50px;
        padding-left: 30%;
        padding-right: 30%;
        width: 100%;
    }
    button {
        background-color: white;
        color: black;
        border: 4px solid #9e4eca;
        border-radius: 4px;
        font-size: 20px;
    }
</style>
