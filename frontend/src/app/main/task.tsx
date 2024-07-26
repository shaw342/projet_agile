"use client"
import React, { useEffect, useState } from 'react';
import styles from './TaskCard.module.css';

interface DataItem {
  id: number;
  name: number;
  content: string;
  State: string;
}

export default function Task() {
  const [data, setData] = useState<DataItem | null>(null);

  
  useEffect(() => {
    fetch('https://jsonplaceholder.typicode.com/posts/1')
      .then(response => response.json())
      .then(data => setData(data))
      .catch(error => console.error(error));
  }, []);

  if (!data) {
    return <div>Loading...</div>;
  }

  return (
    <div className={styles.taskCard}>
    <h2>tache du jour</h2>
    <h3>
      {data.State}
      
    </h3>
  </div>
  );
}