"use client";

import { ITable, ITableHeader, ITableRow } from "@/app/lib/definitions";
import styles from "./styles.module.css";
import { useRouter } from "next/navigation";
import clsx from "clsx";

export default function Table(props: ITable) {
  const router = useRouter();
  return (
    <div className={styles.table_container}>
      <table className={styles.table}>
        <thead>
          <tr className={styles.table_header_row}>
            {props.headers.map((header: ITableHeader, index: number) => (
              <th key={index} className={styles.table_header}>
                {header.title}
              </th>
            ))}
          </tr>
        </thead>
        <tbody className={styles.table_body}>
          {props.rows.map((row: ITableRow, rowIndex: number) => (
            <tr
              className={clsx(row.clickable && styles.table_row_clickable)}
              key={rowIndex}
              onClick={() => {
                if (!row.clickable) return;

                if (row.href) {
                  router.push(row.href);
                }
              }}
            >
              {row.content.map((rData, colIndex) => (
                <td key={colIndex} className={styles.table_body_cell}>
                  {rData}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
