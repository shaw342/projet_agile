"use client"
import React, { useEffect, useState } from 'react';
import styles from './TaskCard.module.css';
export default function Task() {
  useEffect(() =>{
    fetch("http://localhost:8080/api/get/task",{
      headers:{
        "Content-type":"application/json"
      },
      body:JSON.stringify(playload)
    })
  })
  return(
      <div className={styles.taskCard}>
          <h2>tache du jour</h2>
          <p>{task}</p>
      </div>
    )
}