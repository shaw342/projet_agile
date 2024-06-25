"use client";
import "./Register.css";
import { IoMdMail } from "react-icons/io";
import { FormEvent } from "react";
import { IoMail } from "react-icons/io5";

export default function register() {

    async function onSubmit(event: FormEvent<HTMLFormElement>) {
        event.preventDefault();
      
        const form = event.currentTarget;
        const formData = new FormData(form);
      
        // Extract form data and convert to JSON object
        const data = {
          name: formData.get("name"),
          email: formData.get("email"),
          password: formData.get("password"),
        };
      
        const response = await fetch("http://localhost:8080/user", {
          method: "POST",
          credentials: "same-origin",
          body: JSON.stringify(data),
          headers: {
            "Content-Type": "application/json",
          },
        });
      
        const jsonData = await response.json();
        console.log(jsonData);
      }
      
    
    return (<div id="form">
    <div id="title">
        <h1>Signup</h1>
    </div>
    <form onSubmit={onSubmit}>
    <input type="text" name="name" id="name" />
    <IoMail id="mail-icon"/>
    <input type="email" name="email" id="email"/>
    <input type="text" name="password" id="password" />
    <button type="submit"  className="button-3">submit</button>
    </form>
    </div>)
}