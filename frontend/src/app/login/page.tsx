"use client"
import { FormEvent } from "react";
import "./login.css"

export default function login() {
    async function onSubmit(event: FormEvent<HTMLFormElement>) {
        event.preventDefault()
        const form = event.currentTarget;
        const formData = new FormData(form);
        const playload ={
            name:formData.get("name"),
            password:formData.get("password"),
        }
        const response = await fetch("http://localhost:8080/api",{
            method:"POST",
            headers:{
                "Content-Type": "application/json"
            },
            body:JSON.stringify(playload)
        })
        const jsonData = await response.json()
        console.log(jsonData);
        
    }
    return(
        <div id="form">
    <div id="title">
        <h1>login</h1>
    </div>
    <form onSubmit={onSubmit}>
    <input type="text" name="name" id="name" placeholder="name"/>
    <input type="text" name="password" id="password" placeholder="password"/>
    <button type="submit"  className="button-3">submit</button>
    </form>
    </div>
    )
}