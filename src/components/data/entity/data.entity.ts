import {
  Column,
  CreateDateColumn,
  Entity,
  Index,
  JoinColumn,
  ManyToOne,
  PrimaryGeneratedColumn,
  UpdateDateColumn,
} from 'typeorm';
import { User } from '../../user';

@Entity('data')
export default class Data {
  @PrimaryGeneratedColumn({ unsigned: true })
  id!: number;

  @Column({ length: 255 })
  file_url!: string;

  @Column()
  title!: string;

  @Column({ nullable: true, default: null })
  description?: string;

  @Column({ default: false })
  public_yn!: boolean;

  @CreateDateColumn({ type: 'timestamp' })
  created_at!: Date;

  @UpdateDateColumn({ type: 'timestamp' })
  updated_at!: Date;

  @ManyToOne(() => User, (user) => user.datas, {
    createForeignKeyConstraints: false,
  })
  @Index('ix_user_id')
  @JoinColumn({ name: 'user_id', referencedColumnName: 'id' })
  user!: User;
}
