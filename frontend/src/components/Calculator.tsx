"use client"
import React, { useState } from 'react';
import { calculatorClient } from '../services/calculatorClient';
import styles from './Calculator.module.css';

export default function Calculator() {
    // 状态管理
    const [a, setA] = useState('');
    const [b, setB] = useState('');
    const [result, setResult] = useState<number | null>(null);
    const [error, setError] = useState<string | null>(null);
    const [operation, setOperation] = useState<string | null>(null);

    // 处理计算操作
    const handleOperation = async (op: 'add' | 'subtract' | 'multiply' | 'divide') => {
        console.log("用户点击了加法")
        const numA = parseFloat(a);
        const numB = parseFloat(b);

        console.log(`Operation: ${op}, A: ${numA}, B: ${numB}`); // 添加日志记录

        // 输入验证
        if (isNaN(numA) || isNaN(numB)) {
            setError('请输入有效的数字');
            console.error('Invalid input: A or B is not a number'); // 添加错误日志
            return;
        }

        // 除法特殊处理
        if (op === 'divide' && numB === 0) {
            setError('除数不能为零');
            console.error('Division by zero error'); // 添加错误日志
            return;
        }

        setError(null);
        setResult(null);
        setOperation(op);

        try {
            // 调用后端服务
            const response = await calculatorClient[op]({
                a: numA,
                b: numB,
            });
            console.log(`Result: ${response.value}`); // 添加结果日志
            setResult(response.value);
        } catch (err: any) {
            setError(err.message || '计算出错');
            console.error('Error during calculation:', err); // 添加错误日志
        }
    };



    // 获取操作符号
    const getOperationSymbol = () => {
        switch (operation) {
            case 'add': return '+';
            case 'subtract': return '-';
            case 'multiply': return '×';
            case 'divide': return '÷';
            default: return '';
        }
    };

    return (
        <div className={styles.calculator}>
            <h1>使用Go语言调用grpc实现计数器</h1>
            
            <div className={styles.inputGroup}>
                <input
                    type="number"
                    value={a}
                    onChange={(e) => setA(e.target.value)}
                    placeholder="第一个数字"
                    className={styles.input}
                />
                <input
                    type="number"
                    value={b}
                    onChange={(e) => setB(e.target.value)}
                    placeholder="第二个数字"
                    className={styles.input}
                />
            </div>
            
            <div className={styles.buttonGroup}>
                <button onClick={() => handleOperation('add')} className={styles.button}>加法 (+)</button>
                <button onClick={() => handleOperation('subtract')} className={styles.button}>减法 (-)</button>
                <button onClick={() => handleOperation('multiply')} className={styles.button}>乘法 (×)</button>
                <button onClick={() => handleOperation('divide')} className={styles.button}>除法 (÷)</button>

            </div>
            
            {error && <div className={styles.error}>{error}</div>}
            
            {result !== null && (
                <div className={styles.result}>
                    <p>计算式: {a} {getOperationSymbol()} {b} = {result}</p>
                    <p>结果: {result}</p>
                </div>
            )}
        </div>
    );
}
