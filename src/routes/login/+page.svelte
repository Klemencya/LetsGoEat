<script lang="ts">
    import {form, field} from 'svelte-forms';
    import {required} from 'svelte-forms/validators';
    import {base} from "$app/paths";

    const login = field('login', '', [required()]);
    const password = field('password', '', [required()]);

    let myForm = form(login, password);
    let currentUser = '';
    let loginStatus = false;
    let nextLink = 'login'

    async function logInUser() {
        let API_URL = 'http://localhost:8080/api/login'
        myForm = form(login, password)
        const fetchResponse = fetch(API_URL, {
            method: 'POST',
            body: JSON.stringify({
                login: $login.value,
                password: $password.value
            })
        });
        fetchResponse
            .then(response => response.json())
            .then(data => {
                loginStatus = data.msg == "OK";
                if (loginStatus) {
                    currentUser = $login.value;
                    nextLink = "account?user=" + currentUser
                } else {
                    nextLink = "login"
                    alert('Error: No such user!')
                }
            });
    }

</script>

<head>
    <title>Login</title>
    <meta property="og:title" content="Log in"/>
    <meta property="og:description"
          content="Log in and find company to eat."/>
</head>

<section class="form">
    <p><b>Your login:</b> <input type="text" bind:value={$login.value}/></p>
    <p><b>Your password:</b> <input type="text" bind:value={$password.value}/></p>


    {#if loginStatus}
        <a href="{base}/{nextLink}">
            <button>Continue</button>
        </a>
    {:else}
        <button on:click={()=>logInUser()}>Log in</button>
    {/if}

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
