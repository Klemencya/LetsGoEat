<script lang="ts">
    import {base} from "$app/paths";

    let name = '';
    let surname = '';
    let email = '';
    let login = '';
    let password = '';

    async function registerUser() {
        let API_URL = 'http://localhost:8080/api/registration'

        let response = await fetch(API_URL, {
            method: 'POST',
            body: JSON.stringify({
                name,
                surname,
                email,
                login,
                password
            })
        });
        const json = await response.json()
        console.log(json)
    }
</script>

<section class="form">
    <p><b>Your name:</b> <input type="text" bind:value={name} /></p>
    <p><b>Your surname:</b> <input type="text" bind:value={surname} /></p>
    <p><b>Your email:</b> <input type="text" bind:value={email} /></p>
    <p><b>Your login:</b> <input type="text" bind:value={login} /></p>
    <p><b>Your password:</b> <input type="text" bind:value={password} /></p>

    <a href="{base}/account"><button on:click={() => registerUser()}>Create Account</button></a>
</section>

<style>
    section {
        padding-top: 100px;
        display: flex;
        flex-direction: column;
    }

    .form {
        text-align: left;
        margin-right: 30%;
        margin-left: 30%;
        margin-top: 10%;
        border: 2px solid #9e4eca;
        padding: 50px
    }

    p {
        width: 100%;
    }

    input {
        margin-top: 20px;
        padding: 10px;
        width: 100%;
    }
    button {
        margin-top: 10px;
        width: 50%;
        padding: 10px
    }
</style>
