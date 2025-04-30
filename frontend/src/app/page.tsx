import Image from "next/image";
import pageStyles from "./page.module.css";

import Calculator from '../components/Calculator';

export default function Home() {
  return (
    <main className={pageStyles.page}>
      <h1></h1>
      <Calculator />
    </main>
  );
}
