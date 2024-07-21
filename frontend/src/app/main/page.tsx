"use client"
import  Sidebar  from "./sidebar";
import { useEffect,useState } from "react";
import styles from './TaskCard.module.css';
import Task from "./task";

import "./Page.css"
export default function Main() {
    const [data,setData] = useState(null);

  useEffect( () =>{
    const FetchData = async () =>{
      const reponse = await fetch("http://localhost:8080/api/v1/task/get")
      const result = await reponse.json()
      setData(result)
    };
    FetchData();
  },[])
  
    return (
      <>
        <div className="container">
        <Sidebar />
        <main>
        <div className={styles.taskCard}>
          <Task/>
      </div>
        </main> 
    </div>
    </>
    )
}
