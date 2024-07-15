// components/TaskCard.js
import React, { useEffect, useState } from 'react';
import styles from './TaskCard.module.css';
export default function Task() {
  const [data,setData] = useState(null);

  useEffect( () =>{
    const FetchData = async () =>{
      const reponse = await fetch("http://localhost:8080/api/getUser")
      const result = await reponse.json()
      setData(result)
    };
    FetchData();
  },[])
  
  return(
      <div className={styles.taskCard}>
          <h2>tache du jour</h2>
          {data && <pre>{JSON.stringify(data, null, 2)}</pre>}
      </div>
    )
}