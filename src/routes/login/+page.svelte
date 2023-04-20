<script lang="ts">
    import { form, field } from 'svelte-forms';
    import { required } from 'svelte-forms/validators';
    import {base} from "$app/paths";

    let login = '';
    let password= '';
    // const myForm = form(login, password);

    function formIsValid() {
        if (login.length > 0 && password.length > 8){
            return ''
        }
        return 'disabled'
    }
    async function logInUser() {
        let API_URL = 'http://localhost:8080/api/login'

        let response = await fetch(API_URL, {
            method: 'POST',
            body: JSON.stringify({
                login,
                password
            })
        });
        const json = await response.json()
        console.log(json)
    }

</script>

<section class="form">
    <p><b>Your login:</b> <input type="text" bind:value={login} /></p>
    <p><b>Your password:</b> <input type="text" bind:value={password} /></p>

<!--     TODO: разобраться с disabled-->
    <a href="{base}/account"><button on:click={() => logInUser()}>Log in</button></a>
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
