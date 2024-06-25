"use client";
import "./Register.css";
import { FormEvent } from "react";
import { IoMail } from "react-icons/io5";

export default function register() {

    async function onSubmit(event: FormEvent<HTMLFormElement>) {
        event.preventDefault();
      
        const form = event.currentTarget;
        const formData = new FormData(form);
        const data = {
          name:formData.get("name"),
          email:formData.get("email"),
          password:formData.get("password"),

        }
        
        
        // Extract form data and convert to JSON object
      
        const response = await fetch("http://localhost:8080/user", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(data),
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