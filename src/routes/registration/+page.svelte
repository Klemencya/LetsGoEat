<script lang="ts">
    import {base} from "$app/paths";
    import {field, form} from "svelte-forms";
    import {email, min, required} from "svelte-forms/validators";

    let name = field('name', '', [required()]);
    let surname = field('surname', '', [required()]);
    let user_email = field('email', '', [required(), email()]);
    let login = field('login', '', [required()]);
    let password = field('password', '', [required(), min(8)]);
    let myForm = form(name, surname, user_email, login, password);
    let currentUser = '';
    let nextLink = 'registration';
    let registerStatus = false;

    async function registerUser() {
        let API_URL = 'http://localhost:8080/api/registration'
        myForm = form(name, surname, user_email, login, password);

        let fetchResponse = fetch(API_URL, {
            method: 'POST',
            body: JSON.stringify({
                name: $name.value,
                surname: $surname.value,
                email: $user_email.value,
                login: $login.value,
                password: $password.value
            })
        });
        fetchResponse
            .then(response => response.json())
            .then(data => {
                registerStatus = data.msg == "OK";
                if (registerStatus) {
                    currentUser = $login.value;
                    nextLink = "account?user=" + currentUser
                } else {
                    nextLink = "registration"
                    alert('Error: User with such login already exists!')
                }
            });
    }
</script>

<head>
    <title>Registration</title>
    <meta name="title" content="Registration"/>
    <meta name="description"
          content="Register and find company to eat."/>
</head>

<section class="form">
    <p><b>Your name:</b> <input type="text" bind:value={$name.value}/></p>
    <p><b>Your surname:</b> <input type="text" bind:value={$surname.value}/></p>
    <p><b>Your email:</b> <input type="text" bind:value={$user_email.value}/></p>
    <p><b>Your login:</b> <input type="text" bind:value={$login.value}/></p>
    <p><b>Your password:</b> <input type="text" bind:value={$password.value}/></p>

    {#if $myForm.valid}
        {#if registerStatus}
            <a href="{base}/{nextLink}">
                <button>Continue</button>
            </a>
        {:else }
            <button on:click={()=>registerUser()}>Register</button>
        {/if}
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
